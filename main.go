package main

import (
	"main/api"
	"main/config"
	"main/logging"
	"main/util"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var cfg *config.Config

func main() {
	setup()
	log.Infof("%s", PrintVersion())
	for _, line := range cfg.PrintProperties() {
		log.Debugf("Applicazione configurata con proprietà: %s", line)
	}
	for _, line := range util.PrintEnvVars() {
		log.Debugf("Applicazione configurata con variabili di ambiente: %s", line)
	}
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
