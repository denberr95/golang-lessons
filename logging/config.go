package logging

import (
	"main/config"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var cfg *config.LoggingConfig = &config.GetConfig().GoApp.Logging

func Init() {
	log.ReportCaller = cfg.Base.ReportCaller
	configureFormatter()
	configureLogLevel()
}

func GetLogger() *logrus.Logger {
	return log
}
