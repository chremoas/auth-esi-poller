package main

import (
	"github.com/chremoas/services-common/config"
	"github.com/chremoas/auth-srv/proto"
	"github.com/micro/go-micro"
	"github.com/chremoas/esi-srv/proto"
	"fmt"
	"github.com/chremoas/auth-esi-poller/poller"
)

var version = "1.0.0"
var service micro.Service
var name = "auth-esi-poller"

func main() {
	service = config.NewService(version, "poller", name, initialize)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func initialize(configuration *config.Configuration) error {
	entityQueryClient := abaeve_auth.NewEntityQueryService(configuration.LookupService("srv", "auth"), service.Client())
	entityAdminClient := abaeve_auth.NewEntityAdminService(configuration.LookupService("srv", "auth"), service.Client())
	authHandlerClient := abaeve_auth.NewUserAuthenticationService(configuration.LookupService("srv", "auth"), service.Client())
	allianceClient := chremoas_esi.NewAllianceService(configuration.LookupService("srv", "esi"), service.Client())
	corporationClient := chremoas_esi.NewCorporationService(configuration.LookupService("srv", "esi"), service.Client())
	characterClient := chremoas_esi.NewCharacterService(configuration.LookupService("srv", "esi"), service.Client())

	runner := poller.NewAuthEsiPoller(entityQueryClient, entityAdminClient, authHandlerClient, allianceClient, corporationClient, characterClient)

	runner.Start()

	return nil
}
