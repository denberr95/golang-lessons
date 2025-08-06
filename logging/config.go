package logging

import (
	"main/config"
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Init(cfg *config.LoggingConfig) {
	log.ReportCaller = cfg.ReportCaller
	configureFormatter(cfg)
	configureLogLevel(cfg)
}

func configureFormatter(cfg *config.LoggingConfig) {
	switch cfg.Format {
	case config.FormatJSON:
		log.SetFormatter(&logrus.JSONFormatter{
			DisableTimestamp:  cfg.DisableTimestamp,
			DisableHTMLEscape: cfg.DisableHTMLEscaping,
			PrettyPrint:       cfg.PrettyPrint,
			TimestampFormat:   time.RFC3339,
		})
	case config.FormatText:
		log.SetFormatter(&logrus.TextFormatter{
			DisableColors:   cfg.DisableColors,
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
			ForceQuote:      true,
		})
	}
}

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

func Logger() *logrus.Logger {
	return log
}
