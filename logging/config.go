package logging

import (
	"main/config"
	"main/util"

	"github.com/sirupsen/logrus"
)

var defaultLogger = logrus.New()
var loggers = make(map[string]*logrus.Logger)
var loggingConfig *config.LoggingConfig = &config.GetConfig().GoApp.Logging

func Init() {
	defaultLogger.ReportCaller = loggingConfig.Base.ReportCaller
	configureFormatter(defaultLogger)
	configureLogLevel(defaultLogger, loggingConfig.Loggers[util.Root])

	for name, level := range loggingConfig.Loggers {
		if name == util.Root {
			continue
		}

		l := logrus.New()
		l.ReportCaller = loggingConfig.Base.ReportCaller
		configureFormatter(l)
		configureLogLevel(l, level)
		loggers[name] = l
	}
}

func GetLogger() *logrus.Logger {
	return defaultLogger
}

func GetNamedLogger(name string) *logrus.Logger {
	if l, ok := loggers[name]; ok {
		return l
	}
	return defaultLogger
}
