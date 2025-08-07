package logging

import (
	"main/config"

	"github.com/sirupsen/logrus"
)

func configureLogLevel(cfg *config.LoggingConfig) {
	switch cfg.Level {
	case config.PANIC:
		log.SetLevel(logrus.PanicLevel)
	case config.FATAL:
		log.SetLevel(logrus.FatalLevel)
	case config.ERROR:
		log.SetLevel(logrus.ErrorLevel)
	case config.WARN:
		log.SetLevel(logrus.WarnLevel)
	case config.INFO:
		log.SetLevel(logrus.InfoLevel)
	case config.DEBUG:
		log.SetLevel(logrus.DebugLevel)
	case config.TRACE:
		log.SetLevel(logrus.TraceLevel)
	}
}
