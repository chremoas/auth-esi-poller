package main

import (
	"github.com/abaeve/services-common/config"
)

var version string = "1.0.0"

func main() {
	service := config.NewService(version, "auth-esi-poller", initialize)

	service.Run()
}

func initialize(configuration *config.Configuration) error {
	return nil
}
