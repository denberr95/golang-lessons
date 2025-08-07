package main

import (
	"main/api"
	"main/config"
	"main/logging"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var cfg *config.Config
var flags *config.ProgramFlags

func main() {
	setup()
	log.Infof("Application information: %+v", FullVersion())
	log.Debugf("Application started with configuration: %+v", &cfg)
	run()
}

func setup() {
	flags = config.ParseFlags()
	config.Init(flags)
	cfg = config.GetConfig()
	logging.Init(&cfg.GoApp.Logging)
	log = logging.GetLogger()

}

func run() {
	api.Init(&cfg.GoApp.WebServer)
}
