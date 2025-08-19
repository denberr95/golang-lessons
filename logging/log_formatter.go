package logging

import (
	"time"

	"github.com/sirupsen/logrus"
)

func configureFormatter(l *logrus.Logger) {
	if loggingConfig.Text != nil {
		l.SetFormatter(&logrus.TextFormatter{
			DisableColors:   loggingConfig.Text.DisableColors,
			FullTimestamp:   loggingConfig.Text.FullTimestamp,
			ForceQuote:      loggingConfig.Text.ForceQuote,
			TimestampFormat: time.RFC3339,
		})
		return
	}

	if loggingConfig.JSON != nil {
		l.SetFormatter(&logrus.JSONFormatter{
			DisableTimestamp:  loggingConfig.JSON.DisableTimestamp,
			DisableHTMLEscape: loggingConfig.JSON.DisableHTMLEscape,
			PrettyPrint:       loggingConfig.JSON.PrettyPrint,
			TimestampFormat:   time.RFC3339,
		})
		return
	}
}
