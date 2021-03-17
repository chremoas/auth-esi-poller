package main

import (
	"fmt"

	"github.com/chremoas/auth-srv/proto"
	"github.com/chremoas/esi-srv/proto"
	"github.com/chremoas/services-common/config"
	chremoasPrometheus "github.com/chremoas/services-common/prometheus"
	"github.com/micro/go-micro"
	"go.uber.org/zap"

	"github.com/chremoas/auth-esi-poller/poller"
)

var (
	version = "1.0.0"
	service micro.Service
	name    = "auth-esi-poller"
	logger  *zap.Logger
)

func main() {
	var err error
	service = config.NewService(version, "poller", name, initialize)

	// TODO pick stuff up from the config
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Info("Initialized logger")

	go chremoasPrometheus.PrometheusExporter(logger)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func initialize(configuration *config.Configuration) error {
	entityQueryClient := abaeve_auth.EntityQueryServiceClient(configuration.LookupService("srv", "auth"), service.Client())
	entityAdminClient := abaeve_auth.EntityAdminServiceClient(configuration.LookupService("srv", "auth"), service.Client())
	authHandlerClient := abaeve_auth.UserAuthenticationServiceClient(configuration.LookupService("srv", "auth"), service.Client())
	allianceClient := chremoas_esi.NewAllianceService(configuration.LookupService("srv", "esi"), service.Client())
	corporationClient := chremoas_esi.NewCorporationService(configuration.LookupService("srv", "esi"), service.Client())
	characterClient := chremoas_esi.NewCharacterService(configuration.LookupService("srv", "esi"), service.Client())

	runner := poller.NewAuthEsiPoller(entityQueryClient, entityAdminClient, authHandlerClient, allianceClient, corporationClient, characterClient)

	runner.Start()

	return nil
}
