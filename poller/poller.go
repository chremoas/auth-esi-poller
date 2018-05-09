package poller

import (
	"errors"
	"github.com/chremoas/auth-srv/proto"
	"github.com/chremoas/esi-srv/proto"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"time"
)

type AuthEsiPoller interface {
	Start()
	Poll() error
	Stop()
}

type authEsiPoller struct {
	entityQueryClient abaeve_auth.EntityQueryService
	entityAdminClient abaeve_auth.EntityAdminService
	authHanderClient  abaeve_auth.UserAuthenticationService

	allianceClient    chremoas_esi.AllianceService
	corporationClient chremoas_esi.CorporationService
	characterClient   chremoas_esi.CharacterService

	tickTime time.Duration
	ticker   *time.Ticker

	authAllianceMap    map[int32]*abaeve_auth.Alliance
	authCorporationMap map[int32]*abaeve_auth.Corporation

	esiAllianceMap    map[int64]*chremoas_esi.Alliance
	esiCorporationMap map[int64]*chremoas_esi.Corporation

	logger *zap.Logger
}

func (aep *authEsiPoller) Start() {
	aep.ticker = time.NewTicker(aep.tickTime)
	sugar := aep.logger.Sugar()

	sugar.Info("Starting polling loop")
	go func() {
		err := aep.Poll()
		if err != nil {
			//TODO: Replace with logger object
			sugar.Errorf("Received an error while polling: %s", err)
		}
		for range aep.ticker.C {
			err := aep.Poll()
			if err != nil {
				//TODO: Replace with logger object
				sugar.Errorf("Received an error while polling: %s", err)
			}
		}
	}()
}

// Poll currently starts at alliances and works it's way down to characters.  It then walks back up at the corporation
// level and character level if alliance/corporation membership has changed from the last poll.
func (aep *authEsiPoller) Poll() error {
	sugar := aep.logger.Sugar()

	allErrors := ""

	sugar.Info("Calling updateOrDeleteAlliances()")
	err := aep.updateOrDeleteAlliances()
	if err != nil {
		allErrors += err.Error() + "\n"
	}

	sugar.Info("Calling updateOrDeleteCorporations()")
	err = aep.updateOrDeleteCorporations()
	if err != nil {
		allErrors += err.Error() + "\n"
	}

	sugar.Info("Calling updateOrDeleteCharacters()")
	err = aep.updateOrDeleteCharacters()
	if err != nil {
		allErrors += err.Error() + "\n"
	}

	sugar.Info("Calling clearMaps()")
	aep.clearMaps()

	sugar.Info("Syncing to Role service")
	// This needs to be run twice for when people change corps/alliances. Need to figure out why.
	// Some other time. -brian
	_, err = aep.authHanderClient.SyncToRoleService(context.Background(), &abaeve_auth.SyncRequest{})
	if err != nil {
		allErrors += err.Error() + "\n"
	}
	_, err = aep.authHanderClient.SyncToRoleService(context.Background(), &abaeve_auth.SyncRequest{})
	if err != nil {
		allErrors += err.Error() + "\n"
	}

	if len(allErrors) > 0 {
		return errors.New(allErrors)
	}

	return nil
}

func (aep *authEsiPoller) updateOrDeleteAlliances() error {
	sugar := aep.logger.Sugar()

	alliances, err := aep.entityQueryClient.GetAlliances(context.Background(), &abaeve_auth.EntityQueryRequest{EntityType: abaeve_auth.EntityType_ALLIANCE})
	if err != nil {
		return err
	}

	aep.buildAllianceMap(alliances.GetList())

	for _, alliance := range alliances.GetList() {
		response, err := aep.allianceClient.GetAllianceById(context.Background(), &chremoas_esi.GetAllianceByIdRequest{Id: int32(alliance.Id)})
		if err == nil {
			if response.Alliance == nil {
				sugar.Infof("Removing alliance: %s", alliance)
				aep.entityAdminClient.AllianceUpdate(context.Background(), &abaeve_auth.AllianceAdminRequest{
					Alliance:  alliance,
					Operation: abaeve_auth.EntityOperation_REMOVE,
				})
			} else if allianceDiffers(alliance, response.Alliance) {
				aep.authAllianceMap[int32(alliance.Id)] = &abaeve_auth.Alliance{
					Id:     alliance.Id,
					Name:   response.Alliance.Name,
					Ticker: response.Alliance.Ticker,
				}

				sugar.Infof("Updating alliance: %s", aep.authAllianceMap[int32(alliance.Id)])
				aep.entityAdminClient.AllianceUpdate(context.Background(), &abaeve_auth.AllianceAdminRequest{
					Alliance:  aep.authAllianceMap[int32(alliance.Id)],
					Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
				})
			}

			aep.esiAllianceMap[alliance.Id] = response.Alliance
		} else {
			//TODO: Do stuff with this error
		}
	}

	return nil
}

func (aep *authEsiPoller) updateOrDeleteCorporations() error {
	sugar := aep.logger.Sugar()

	corporations, err := aep.entityQueryClient.GetCorporations(context.Background(), &abaeve_auth.EntityQueryRequest{EntityType: abaeve_auth.EntityType_CORPORATION})
	if err != nil {
		return err
	}

	aep.buildCorporationMap(corporations.GetList())

	for _, corporation := range corporations.GetList() {
		response, err := aep.corporationClient.GetCorporationById(context.Background(), &chremoas_esi.GetCorporationByIdRequest{Id: int32(corporation.Id)})
		if err == nil {
			if response.Corporation == nil {
				sugar.Infof("Removing corporation: %s", corporation)
				aep.entityAdminClient.CorporationUpdate(context.Background(), &abaeve_auth.CorporationAdminRequest{
					Corporation: corporation,
					Operation:   abaeve_auth.EntityOperation_REMOVE,
				})
			} else if corporationDiffers(corporation, response.Corporation) {
				aep.checkAndUpdateCorpsAllianceIfNecessary(corporation, response.Corporation)

				sugar.Infof("Updating corporation: %s", corporation)
				aep.entityAdminClient.CorporationUpdate(context.Background(), &abaeve_auth.CorporationAdminRequest{
					Corporation: &abaeve_auth.Corporation{
						Id:         corporation.Id,
						Name:       response.Corporation.Name,
						Ticker:     response.Corporation.Ticker,
						AllianceId: int64(response.Corporation.AllianceId),
					},
					Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
				})
			}
		} else {
			//TODO: Do stuff with this error
		}
	}

	return nil
}

func (aep *authEsiPoller) updateOrDeleteCharacters() error {
	sugar := aep.logger.Sugar()

	characters, err := aep.entityQueryClient.GetCharacters(context.Background(), &abaeve_auth.EntityQueryRequest{EntityType: abaeve_auth.EntityType_CHARACTER})
	if err != nil {
		return err
	}

	allNonFatalErrors := ""

	for _, character := range characters.GetList() {
		response, err := aep.characterClient.GetCharacterById(context.Background(), &chremoas_esi.GetCharacterByIdRequest{Id: int32(character.Id)})
		if err == nil {
			if response.Character == nil {
				sugar.Infof("Removing character: %s", character)
				aep.entityAdminClient.CharacterUpdate(context.Background(), &abaeve_auth.CharacterAdminRequest{
					Character: character,
					Operation: abaeve_auth.EntityOperation_REMOVE,
				})
			} else if characterDiffers(character, response.Character) {
				if character.CorporationId != int64(response.Character.CorporationId) && aep.esiCorporationMap[character.CorporationId] == nil {
					esiResponse, err := aep.corporationClient.GetCorporationById(context.Background(), &chremoas_esi.GetCorporationByIdRequest{
						Id: response.Character.CorporationId,
					})
					if err != nil {
						allNonFatalErrors += err.Error() + "\n"
					} else {
						aep.checkAndUpdateCorpsAllianceIfNecessary(aep.authCorporationMap[int32(character.CorporationId)], esiResponse.Corporation)

						newAuthCorporation := &abaeve_auth.Corporation{
							Id:         int64(response.Character.CorporationId),
							Name:       esiResponse.Corporation.Name,
							Ticker:     esiResponse.Corporation.Ticker,
							AllianceId: int64(esiResponse.Corporation.AllianceId),
						}

						sugar.Infof("Updating character corporation: %s", character)
						aep.entityAdminClient.CorporationUpdate(context.Background(), &abaeve_auth.CorporationAdminRequest{
							Corporation: newAuthCorporation,
							Operation:   abaeve_auth.EntityOperation_ADD_OR_UPDATE,
						})

						aep.esiCorporationMap[character.CorporationId] = esiResponse.Corporation
						aep.authCorporationMap[int32(character.CorporationId)] = newAuthCorporation
					}
				}

				sugar.Infof("Updating character: %s", character)
				aep.entityAdminClient.CharacterUpdate(context.Background(), &abaeve_auth.CharacterAdminRequest{
					Character: &abaeve_auth.Character{
						Id:            character.Id,
						Name:          response.Character.Name,
						CorporationId: int64(response.Character.CorporationId),
					},
				})
			}
		} else {
			allNonFatalErrors += err.Error() + "\n"
		}
	}

	if len(allNonFatalErrors) > 0 {
		return errors.New(allNonFatalErrors)
	}

	return nil
}

func (aep *authEsiPoller) checkAndUpdateCorpsAllianceIfNecessary(authCorporation *abaeve_auth.Corporation, esiCorporation *chremoas_esi.Corporation) error {
	sugar := aep.logger.Sugar()

	if esiCorporation.AllianceId == 0 {
		return nil
	}

	sugar.Infof("Updating corporations alliance for %s with allianceId %d\n", esiCorporation.Name, esiCorporation.AllianceId)
	allErrors := ""

	if authCorporation.AllianceId != int64(esiCorporation.AllianceId) && aep.esiAllianceMap[int64(esiCorporation.AllianceId)] == nil {
		newAllianceResponse, err := aep.allianceClient.GetAllianceById(context.Background(), &chremoas_esi.GetAllianceByIdRequest{
			Id: esiCorporation.AllianceId,
		})
		if err != nil {
			allErrors += err.Error() + "\n"
		}

		aep.authAllianceMap[esiCorporation.AllianceId] = &abaeve_auth.Alliance{
			Id:     int64(esiCorporation.AllianceId),
			Name:   newAllianceResponse.Alliance.Name,
			Ticker: newAllianceResponse.Alliance.Ticker,
		}

		aep.esiAllianceMap[int64(esiCorporation.AllianceId)] = newAllianceResponse.Alliance

		sugar.Infof("Updating alliance: %s", aep.authAllianceMap[esiCorporation.AllianceId])
		_, err = aep.entityAdminClient.AllianceUpdate(context.Background(), &abaeve_auth.AllianceAdminRequest{
			Alliance:  aep.authAllianceMap[esiCorporation.AllianceId],
			Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
		})
		if err != nil {
			allErrors += err.Error() + "\n"
		}
	}

	if len(allErrors) > 0 {
		return errors.New(allErrors)
	}

	return nil
}

func (aep *authEsiPoller) buildAllianceMap(alliances []*abaeve_auth.Alliance) {
	if aep.authAllianceMap == nil {
		aep.authAllianceMap = make(map[int32]*abaeve_auth.Alliance)
	}

	for _, alliance := range alliances {
		aep.authAllianceMap[int32(alliance.Id)] = alliance
	}
}

func (aep *authEsiPoller) buildCorporationMap(corporations []*abaeve_auth.Corporation) {
	if aep.authCorporationMap == nil {
		aep.authCorporationMap = make(map[int32]*abaeve_auth.Corporation)
	}

	for _, corporation := range corporations {
		aep.authCorporationMap[int32(corporation.Id)] = corporation
	}
}

func (aep *authEsiPoller) clearMaps() {
	aep.authAllianceMap = make(map[int32]*abaeve_auth.Alliance)
	aep.authCorporationMap = make(map[int32]*abaeve_auth.Corporation)
}

func allianceDiffers(authAlliance *abaeve_auth.Alliance, esiAlliance *chremoas_esi.Alliance) bool {
	if authAlliance.Name != esiAlliance.Name || authAlliance.Ticker != esiAlliance.Ticker {
		return true
	}
	return false
}

func corporationDiffers(authCorporation *abaeve_auth.Corporation, esiCorporation *chremoas_esi.Corporation) bool {
	if authCorporation.Name != esiCorporation.Name || authCorporation.Ticker != esiCorporation.Ticker || authCorporation.AllianceId != int64(esiCorporation.AllianceId) {
		return true
	}
	return false
}

func characterDiffers(authCharacter *abaeve_auth.Character, esiCharacter *chremoas_esi.Character) bool {
	if authCharacter.Name != esiCharacter.Name || authCharacter.CorporationId != int64(esiCharacter.CorporationId) {
		return true
	}
	return false
}

func (aep *authEsiPoller) Stop() {
	aep.ticker.Stop()
}

func NewAuthEsiPoller(eqc abaeve_auth.EntityQueryService,
	eac abaeve_auth.EntityAdminService,
	ahc abaeve_auth.UserAuthenticationService,
	allianceClient chremoas_esi.AllianceService,
	corporationClient chremoas_esi.CorporationService,
	characterClient chremoas_esi.CharacterService) AuthEsiPoller {

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Info("Setting up Auth ESI Poller")

	return &authEsiPoller{
		entityAdminClient: eac,
		entityQueryClient: eqc,
		authHanderClient:  ahc,
		allianceClient:    allianceClient,
		corporationClient: corporationClient,
		characterClient:   characterClient,
		tickTime:          time.Minute * 5,

		authAllianceMap:    make(map[int32]*abaeve_auth.Alliance),
		authCorporationMap: make(map[int32]*abaeve_auth.Corporation),

		esiAllianceMap:    make(map[int64]*chremoas_esi.Alliance),
		esiCorporationMap: make(map[int64]*chremoas_esi.Corporation),
		logger:            logger,
	}
}
