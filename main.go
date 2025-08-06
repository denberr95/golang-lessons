package main

import (
	"log"
	"main/config"
	"main/logging"
)

func main() {
	flags := config.ParseFlags()
	cfg, err := config.LoadConfig(flags)
	if err != nil {
		log.Fatalf("%v", err)
	}
	logging.Init(&cfg.GoApp.Logging)
	logging.Logger().Infof("application started with configuration: %+v", *cfg)
}
