package main

import (
	"main/config"
	"main/logging"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var cfg *config.Config
var flags *config.ProgramFlags

func main() {
	initialize()
	log.Infof("starting application: %+v", FullVersion())
	log.Debugf("application started with configuration: %+v", &cfg)
}

func initialize() {
	flags = config.ParseFlags()
	config.Init(flags)
	cfg = config.GetConfig()
	logging.Init(&cfg.GoApp.Logging)
	log = logging.GetLogger()
}
