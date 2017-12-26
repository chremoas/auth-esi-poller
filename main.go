package main

import (
	"github.com/chremoas/services-common/config"
	"github.com/chremoas/auth-srv/proto"
	"github.com/micro/go-micro"
	"github.com/chremoas/esi-srv/proto"
	"fmt"
	"github.com/chremoas/auth-esi-poller/poller"
)

var version string = "1.0.0"
var service micro.Service
var name = "auth-esi-poller"

func main() {
	service = config.NewService(version, "poller", name, initialize)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func initialize(configuration *config.Configuration) error {
	entityQueryClient := abaeve_auth.NewEntityQueryClient(configuration.LookupService("srv", "auth"), service.Client())
	entityAdminClient := abaeve_auth.NewEntityAdminClient(configuration.LookupService("srv", "auth"), service.Client())
	allianceClient := chremoas_esi.NewAllianceServiceClient(configuration.LookupService("srv", "esi"), service.Client())
	corporationClient := chremoas_esi.NewCorporationServiceClient(configuration.LookupService("srv", "esi"), service.Client())
	characterClient := chremoas_esi.NewCharacterServiceClient(configuration.LookupService("srv", "esi"), service.Client())

	runner := poller.NewAuthEsiPoller(entityQueryClient, entityAdminClient, allianceClient, corporationClient, characterClient)

	runner.Start()

	return nil
}
