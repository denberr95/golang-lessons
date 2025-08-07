package main

import (
	"main/config"
	"main/logging"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func main() {
	flags := config.ParseFlags()
	cfg, err := config.LoadConfig(flags)
	if err != nil {
		log.Fatalf("%v", err)
	}
	logging.Init(&cfg.GoApp.Logging)
	log = logging.Logger()
	log.Infof("starting application: %+v", FullVersion())
	log.Debugf("application started with configuration: %+v", *cfg)
}
