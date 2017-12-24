package poller

import (
	"fmt"
	"github.com/chremoas/auth-srv/proto"
	"github.com/chremoas/esi-srv/proto"
	"time"
	"golang.org/x/net/context"
)

type AuthEsiPoller interface {
	Start()
	Poll() error
	Stop()
}

type authEsiPoller struct {
	entityQueryClient abaeve_auth.EntityQueryClient
	entityAdminClient abaeve_auth.EntityAdminClient

	allianceClient    chremoas_esi.AllianceServiceClient
	corporationClient chremoas_esi.CorporationServiceClient
	characterClient   chremoas_esi.CharacterServiceClient

	tickTime time.Duration
	ticker   *time.Ticker
}

func (aep *authEsiPoller) Start() {
	aep.ticker = time.NewTicker(aep.tickTime)

	go func() {
		for range aep.ticker.C {
			err := aep.Poll()
			if err != nil {
				//TODO: Replace with logger object
				fmt.Printf("Received an error while polling: %s\n", err)
			}
		}
	}()
}

func (aep *authEsiPoller) Poll() error {
	allErrors := ""

	err := aep.updateOrDeleteAlliances()
	if err != nil {
		allErrors += err.Error() + "\n"
	}

	err = aep.updateOrDeleteCorporations()
	if err != nil {
		allErrors += err.Error() + "\n"
	}

	err = aep.updateOrDeleteCharacters()
	if err != nil {
		allErrors += err.Error() + "\n"
	}

	return nil
}

func (aep *authEsiPoller) updateOrDeleteAlliances() error {
	alliances, err := aep.entityQueryClient.GetAlliances(context.Background(), &abaeve_auth.EntityQueryRequest{EntityType: abaeve_auth.EntityType_ALLIANCE})
	if err != nil {
		return err
	}

	for _, alliance := range alliances.GetList() {
		response, err := aep.allianceClient.GetAllianceById(context.Background(), &chremoas_esi.GetAllianceByIdRequest{ Id: int32(alliance.Id) })
		if err == nil {
			if response.Alliance == nil {
				aep.entityAdminClient.AllianceUpdate(context.Background(), &abaeve_auth.AllianceAdminRequest{
					Alliance: alliance,
					Operation: abaeve_auth.EntityOperation_REMOVE,
				})
			} else if allianceDiffers(alliance, response.Alliance) {
				aep.entityAdminClient.AllianceUpdate(context.Background(), &abaeve_auth.AllianceAdminRequest{
					Alliance: &abaeve_auth.Alliance{
						Id: alliance.Id,
						Name: response.Alliance.Name,
						Ticker: response.Alliance.Ticker,
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

func (aep *authEsiPoller) updateOrDeleteCorporations() error {
	corporations, err := aep.entityQueryClient.GetCorporations(context.Background(), &abaeve_auth.EntityQueryRequest{EntityType: abaeve_auth.EntityType_CORPORATION})
	if err != nil {
		return err
	}

	for _, corporation := range corporations.GetList() {
		response, err := aep.corporationClient.GetCorporationById(context.Background(), &chremoas_esi.GetCorporationByIdRequest{ Id: int32(corporation.Id) })
		if err == nil {
			if response.Corporation == nil {
				aep.entityAdminClient.CorporationUpdate(context.Background(), &abaeve_auth.CorporationAdminRequest{
					Corporation: corporation,
					Operation: abaeve_auth.EntityOperation_REMOVE,
				})
			} else if corporationDiffers(corporation, response.Corporation) {
				aep.entityAdminClient.CorporationUpdate(context.Background(), &abaeve_auth.CorporationAdminRequest{
					Corporation: &abaeve_auth.Corporation{
						Id: corporation.Id,
						Name: response.Corporation.Name,
						Ticker: response.Corporation.Ticker,
						AllianceId: int64(response.Corporation.AllianceId),
					},
				})
			}
		} else {
			//TODO: Do stuff with this error
		}
	}

	return nil
}

func (aep *authEsiPoller) updateOrDeleteCharacters() error {
	return nil
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

func NewAuthEsiPoller(eqc abaeve_auth.EntityQueryClient,
	eac abaeve_auth.EntityAdminClient,
	allianceClient chremoas_esi.AllianceServiceClient,
	corporationClient chremoas_esi.CorporationServiceClient,
	characterClient chremoas_esi.CharacterServiceClient) AuthEsiPoller {

	return &authEsiPoller{
		entityAdminClient: eac,
		entityQueryClient: eqc,
		allianceClient:    allianceClient,
		corporationClient: corporationClient,
		characterClient:   characterClient,
		tickTime:          time.Minute * 5,
	}
}
