package logging

import (
	"main/config"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var loggingConfig *config.LoggingConfig = &config.GetConfig().GoApp.Logging

func Init() {
	log.ReportCaller = loggingConfig.Base.ReportCaller
	configureFormatter()
	configureLogLevel()
}

func GetLogger() *logrus.Logger {
	return log
}
