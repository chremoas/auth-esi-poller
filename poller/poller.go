package poller

import (
	"fmt"
	"github.com/chremoas/auth-esi-poller/esiapi"
	"github.com/chremoas/auth-srv/proto"
	"time"
)

type AuthEsiPoller interface {
	Start()
	Poll() error
	Stop()
}

type authEsiPoller struct {
	entityQueryClient abaeve_auth.EntityQueryClient
	entityAdminClient abaeve_auth.EntityAdminClient

	allianceApi    esiapi.AllianceApi
	corporationApi esiapi.CorporationApi
	characterApi   esiapi.CharacterApi

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
	return nil
}

func (aep *authEsiPoller) Stop() {
	aep.ticker.Stop()
}

func NewAuthEsiPoller(eqc abaeve_auth.EntityQueryClient, eac abaeve_auth.EntityAdminClient, alliApi esiapi.AllianceApi, corpApi esiapi.CorporationApi, charApi esiapi.CharacterApi) AuthEsiPoller {
	return &authEsiPoller{
		entityAdminClient: eac,
		entityQueryClient: eqc,
		allianceApi:       alliApi,
		corporationApi:    corpApi,
		characterApi:      charApi,
		tickTime:          time.Minute * 5,
	}
}
