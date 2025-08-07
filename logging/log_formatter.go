package logging

import (
	"main/config"
	"time"

	"github.com/sirupsen/logrus"
)

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
			FullTimestamp:   cfg.FullTimestamp,
			TimestampFormat: time.RFC3339,
			ForceQuote:      cfg.ForceQuote,
		})
	}
}
