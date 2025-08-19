package logging

import (
	"main/config"

	"github.com/sirupsen/logrus"
)

func configureLogLevel(l *logrus.Logger, level config.LogLevel) {
	if level.IsValid() {
		l.SetLevel(convertLogLevel(level))
	}
}

func convertLogLevel(level config.LogLevel) logrus.Level {
	switch level {
	case config.PANIC:
		return logrus.PanicLevel
	case config.FATAL:
		return logrus.FatalLevel
	case config.ERROR:
		return logrus.ErrorLevel
	case config.WARN:
		return logrus.WarnLevel
	case config.INFO:
		return logrus.InfoLevel
	case config.DEBUG:
		return logrus.DebugLevel
	case config.TRACE:
		return logrus.TraceLevel
	default:
		return logrus.ErrorLevel
	}
}
