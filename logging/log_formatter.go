package logging

import (
	"time"

	"github.com/sirupsen/logrus"
)

func configureFormatter() {
	if cfg.Text != nil {
		log.SetFormatter(&logrus.TextFormatter{
			DisableColors:   cfg.Text.DisableColors,
			FullTimestamp:   cfg.Text.FullTimestamp,
			ForceQuote:      cfg.Text.ForceQuote,
			TimestampFormat: time.RFC3339,
		})
		return
	}

	if cfg.JSON != nil {
		log.SetFormatter(&logrus.JSONFormatter{
			DisableTimestamp:  cfg.JSON.DisableTimestamp,
			DisableHTMLEscape: cfg.JSON.DisableHTMLEscape,
			PrettyPrint:       cfg.JSON.PrettyPrint,
			TimestampFormat:   time.RFC3339,
		})
		return
	}
}
