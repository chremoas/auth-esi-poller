package main

import (
	"github.com/chremoas/services-common/config"
)

var version string = "1.0.0"

func main() {
	service := config.NewService(version, "poller", "auth-esi-poller", initialize)

	service.Run()
}

func initialize(configuration *config.Configuration) error {
	return nil
}
