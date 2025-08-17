package logging

import (
	"time"

	"github.com/sirupsen/logrus"
)

func configureFormatter() {
	if loggingConfig.Text != nil {
		log.SetFormatter(&logrus.TextFormatter{
			DisableColors:   loggingConfig.Text.DisableColors,
			FullTimestamp:   loggingConfig.Text.FullTimestamp,
			ForceQuote:      loggingConfig.Text.ForceQuote,
			TimestampFormat: time.RFC3339,
		})
		return
	}

	if loggingConfig.JSON != nil {
		log.SetFormatter(&logrus.JSONFormatter{
			DisableTimestamp:  loggingConfig.JSON.DisableTimestamp,
			DisableHTMLEscape: loggingConfig.JSON.DisableHTMLEscape,
			PrettyPrint:       loggingConfig.JSON.PrettyPrint,
			TimestampFormat:   time.RFC3339,
		})
		return
	}
}
