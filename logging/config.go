package logging

import (
	"main/config"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Init(cfg *config.LoggingConfig) {
	log.ReportCaller = cfg.ReportCaller
	configureFormatter(cfg)
	configureLogLevel(cfg)
}

func Logger() *logrus.Logger {
	return log
}
