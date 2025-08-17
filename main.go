package main

import (
	"main/api"
	"main/config"
	"main/logging"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var cfg *config.Config

func main() {
	setup()
	log.Infof("%s", PrintVersion())
	log.Debugf("Applicazione avviata con: %s", cfg.Print())
	start()
}

func setup() {
	config.Load()
	cfg = config.GetConfig()
	logging.Init()
	log = logging.GetLogger()
}

func start() {
	api.Run()
}
