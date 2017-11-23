package poller

import (
	"github.com/chremoas/auth-esi-poller/mock"
	authsrv_mocks "github.com/chremoas/auth-srv/mocks"
	esi_mocks "github.com/chremoas/esi-srv/proto"
	"github.com/chremoas/auth-srv/proto"
	"github.com/antihax/goesi/esi"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func PollerTestSetup(t *testing.T) (*gomock.Controller,
	*mocks.MockAllianceApi,
	*mocks.MockCorporationApi,
	*mocks.MockCharacterApi,
	*authsrv_mocks.MockEntityQueryClient,
	*authsrv_mocks.MockEntityAdminClient,
	*esi_mocks.MockEntityQueryClient) {

	mockCtrl := gomock.NewController(t)

	mockAllianceApi := mocks.NewMockAllianceApi(mockCtrl)
	mockCorporationApi := mocks.NewMockCorporationApi(mockCtrl)
	mockCharacterApi := mocks.NewMockCharacterApi(mockCtrl)

	mockEntityQueryClient := authsrv_mocks.NewMockEntityQueryClient(mockCtrl)
	mockEntityAdminClient := authsrv_mocks.NewMockEntityAdminClient(mockCtrl)

	mockEsiEntityQueryClient := esi_mocks.NewMockEntityQueryClient(mockCtrl)

	return mockCtrl, mockAllianceApi, mockCorporationApi, mockCharacterApi, mockEntityQueryClient, mockEntityAdminClient, mockEsiEntityQueryClient
}

func TestAuthEsiPoller_Poll_AllianceClosed(t *testing.T) {
	mockCtrl, mockAllianceApi, mockCorporationApi, mockCharacterApi, mockEntityQueryClient, mockEntityAdminClient, mockEsiEntityQueryClient := PollerTestSetup(t)
	defer mockCtrl.Finish()

	poller := NewAuthEsiPoller(mockEntityQueryClient, mockEntityAdminClient, mockAllianceApi, mockCorporationApi, mockCharacterApi)
	poller.Poll()

	//What alliances do we know about?
	mockEntityQueryClient.EXPECT().GetAlliances(gomock.Any(),
		&abaeve_auth.EntityQueryRequest{
			EntityType: abaeve_auth.EntityType_ALLIANCE,
		},
		gomock.Any(),
	).Return(&abaeve_auth.AlliancesResponse{
		List: []*abaeve_auth.Alliance{
			{
				Id:     int64(1),
				Name:   "Alliance Name 1",
				Ticker: "A T 1",
			},
			{
				Id:     int64(2),
				Name:   "Alliance Name 2",
				Ticker: "A T 2",
			},
		},
	}, nil)

	//Do they still exist?
	mockAllianceApi.EXPECT().GetAlliancesAllianceId(int32(1), gomock.Nil()).Times(1).Return(nil, nil, nil)
	mockAllianceApi.EXPECT().GetAlliancesAllianceId(int32(2), gomock.Nil()).Times(1).Return(esi.GetAlliancesAllianceIdOk{
		AllianceName: "Alliance Name 2",
		Ticker:       "A T 2",
		ExecutorCorp: int32(1),
		DateFounded:  time.Now(),
	}, nil, nil)

	//We should be asking the auth-srv to delete the Alliance.
	mockEntityAdminClient.EXPECT().AllianceUpdate(gomock.Any(), &abaeve_auth.AllianceAdminRequest{
		Alliance: &abaeve_auth.Alliance{
			Id: int64(1),
		},
	}, gomock.Any()).Return(&abaeve_auth.EntityAdminResponse{
		Success: true,
	})
}

func TestAuthEsiPoller_Poll_AllianceNoChange(t *testing.T) {
	mockCtrl, mockAllianceApi, mockCorporationApi, mockCharacterApi, mockEntityQueryClient, mockEntityAdminClient, mockEsiEntityQueryClient := PollerTestSetup(t)
	defer mockCtrl.Finish()

	poller := NewAuthEsiPoller(mockEntityQueryClient, mockEntityAdminClient, mockAllianceApi, mockCorporationApi, mockCharacterApi)
	poller.Poll()

	//What alliances do we know about?
	mockEntityQueryClient.EXPECT().GetAlliances(gomock.Any(),
		&abaeve_auth.EntityQueryRequest{
			EntityType: abaeve_auth.EntityType_ALLIANCE,
		},
		gomock.Any(),
	).Return(&abaeve_auth.AlliancesResponse{
		List: []*abaeve_auth.Alliance{
			{
				Id:     int64(1),
				Name:   "Alliance Name 1",
				Ticker: "A T 1",
			},
			{
				Id:     int64(2),
				Name:   "Alliance Name 2",
				Ticker: "A T 2",
			},
		},
	}, nil)

	//Do they still exist?
	mockAllianceApi.EXPECT().GetAlliancesAllianceId(int32(1), gomock.Nil()).Times(1).Return(esi.GetAlliancesAllianceIdOk{
		AllianceName: "Alliance Name 1",
		Ticker:       "A T 1",
		ExecutorCorp: int32(1),
		DateFounded:  time.Now(),
	}, nil, nil)
	mockAllianceApi.EXPECT().GetAlliancesAllianceId(int32(2), gomock.Nil()).Times(1).Return(esi.GetAlliancesAllianceIdOk{
		AllianceName: "Alliance Name 2",
		Ticker:       "A T 2",
		ExecutorCorp: int32(1),
		DateFounded:  time.Now(),
	}, nil, nil)

	//Expect that we don't communicate with the auth-srv any longer
	mockEntityAdminClient.EXPECT().AllianceUpdate(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	mockEntityAdminClient.EXPECT().CorporationUpdate(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	mockEntityAdminClient.EXPECT().CharacterUpdate(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
}

func TestNewAuthEsiPoller_Poll_CorporationMovedAlliance(t *testing.T) {
	mockCtrl, mockAllianceApi, mockCorporationApi, mockCharacterApi, mockEntityQueryClient, mockEntityAdminClient, mockEsiEntityQueryClient := PollerTestSetup(t)
	defer mockCtrl.Finish()

	poller := NewAuthEsiPoller(mockEntityQueryClient, mockEntityAdminClient, mockAllianceApi, mockCorporationApi, mockCharacterApi)
	poller.Poll()

	//What corporations do we have?
	mockEntityQueryClient.EXPECT().GetCorporations(gomock.Any(),
		&abaeve_auth.EntityQueryRequest{
			EntityType: abaeve_auth.EntityType_CORPORATION,
		}, gomock.Any()).Times(1).Return(&abaeve_auth.CorporationsResponse{
		List: []*abaeve_auth.Corporation{
			{
				Id:         int64(1),
				Name:       "Corporation Name 1",
				Ticker:     "C T 1",
				AllianceId: int64(1),
			},
		},
	}, nil)

	//Do they still exist?
	mockCorporationApi.EXPECT().GetCorporationsCorporationId(int32(1), gomock.Nil()).Times(1).Return(
		esi.GetCorporationsCorporationIdOk{
			CorporationName: "Corporation Name 1",
			Ticker:          "C T 1",
			AllianceId:      int32(2),
		},
		nil, nil)

	//Update things in the auth-srv
	mockEntityAdminClient.EXPECT().CorporationUpdate(gomock.Any(), &abaeve_auth.CorporationAdminRequest{
		Corporation: &abaeve_auth.Corporation{
			AllianceId: int64(2),
			Id:         int64(1),
			Ticker:     "C T 1",
			Name:       "Corporation Name 1",
		},
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
	}, gomock.Any()).Times(1).Return(&abaeve_auth.EntityAdminResponse{
		Success: true,
	})
}
func TestAuthEsiPoller_Poll_CorporationClosed(t *testing.T) {
	mockCtrl, mockAllianceApi, mockCorporationApi, mockCharacterApi, mockEntityQueryClient, mockEntityAdminClient, mockEsiEntityQueryClient := PollerTestSetup(t)
	defer mockCtrl.Finish()

	poller := NewAuthEsiPoller(mockEntityQueryClient, mockEntityAdminClient, mockAllianceApi, mockCorporationApi, mockCharacterApi)
	poller.Poll()

	//What corporations do we have?
	mockEntityQueryClient.EXPECT().GetCorporations(gomock.Any(),
		&abaeve_auth.EntityQueryRequest{
			EntityType: abaeve_auth.EntityType_CORPORATION,
		}, gomock.Any()).Times(1).Return(&abaeve_auth.CorporationsResponse{
		List: []*abaeve_auth.Corporation{
			{
				Id:         int64(1),
				Name:       "Corporation Name 1",
				Ticker:     "C T 1",
				AllianceId: int64(1),
			},
			{
				Id:         int64(2),
				Name:       "Corporation Name 2",
				Ticker:     "C T 2",
				AllianceId: int64(2),
			},
		},
	}, nil)

	//Do they still exist?
	mockCorporationApi.EXPECT().GetCorporationsCorporationId(int32(1), gomock.Nil()).Times(1).Return(
		esi.GetCorporationsCorporationIdOk{
			CorporationName: "Corporation Name 1",
			Ticker:          "C T 1",
			AllianceId:      int32(1),
		},
		nil, nil)
	mockCorporationApi.EXPECT().GetCorporationsCorporationId(int32(2), gomock.Nil()).Times(1).Return(nil, nil, nil)

	//Make the appropriate updates in the auth-srv
	mockEntityAdminClient.EXPECT().CorporationUpdate(gomock.Any(), &abaeve_auth.CorporationAdminRequest{
		Corporation: &abaeve_auth.Corporation{
			Id: int64(2),
		},
		Operation: abaeve_auth.EntityOperation_REMOVE,
	}, gomock.Any())
}

func TestAuthEsiPoller_Poll_CharacterLeftCorp(t *testing.T) {
	mockCtrl, mockAllianceApi, mockCorporationApi, mockCharacterApi, mockEntityQueryClient, mockEntityAdminClient, mockEsiEntityQueryClient := PollerTestSetup(t)
	defer mockCtrl.Finish()

	poller := NewAuthEsiPoller(mockEntityQueryClient, mockEntityAdminClient, mockAllianceApi, mockCorporationApi, mockCharacterApi)
	poller.Poll()

	//What characters do we know about?
	mockEntityQueryClient.EXPECT().GetCharacters(gomock.Any(),
		&abaeve_auth.EntityQueryRequest{
			EntityType: abaeve_auth.EntityType_CHARACTER,
		}, gomock.Any()).Times(1).Return(&abaeve_auth.CharactersResponse{
		List: []*abaeve_auth.Character{
			{
				Id:            int64(1),
				Name:          "Character Name 1",
				CorporationId: int64(1),
			},
		},
	}, nil)

	//Do they still exist?
	mockCharacterApi.EXPECT().GetCharactersCharacterId(int64(1), gomock.Any()).Return(
		esi.GetCharactersCharacterIdOk{
			CorporationId: int32(2),
			Name:          "Character Name 1",
		}, nil, nil)

	//Make the appropriate update in the auth-srv
	mockEntityAdminClient.EXPECT().CharacterUpdate(gomock.Any(), &abaeve_auth.CharacterAdminRequest{
		Character: &abaeve_auth.Character{
			Id:            int64(1),
			Name:          "Character Name 1",
			CorporationId: int64(2),
		},
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
	}, gomock.Any()).Times(1).Return(abaeve_auth.EntityAdminResponse{
		Success: true,
	})
}
