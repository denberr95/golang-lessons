package config

import (
	"main/util"
	"net/url"
	"os"
	"strings"
	"time"
)

func applyConfiguration(cfg *Config) {
	validateWebServerConfig(&cfg.GoApp.WebServer)
	validateLoggingConfig(&cfg.GoApp.Logging)
}

func validateWebServerConfig(cfg *WebServerConfig) {
	validateHost(cfg)
	validateHttpPort(cfg)
	validateBasePath(cfg)
	validateReadTimeout(cfg)
	validateWriteTimeout(cfg)
	validateMaxHeaderSizeMB(cfg)
	validateMaxMultipartMemoryMB(cfg)
	validateIdleTimeout(cfg)
	validateGracefulShutdownTime(cfg)
}

func validateLoggingConfig(cfg *LoggingConfig) {
	validateLogFormat(cfg)
	validateLogLevel(cfg.Loggers)
}

func validateLogFormat(cfg *LoggingConfig) {
	if cfg.Text == nil && cfg.JSON == nil {
		cfg.Text = &LoggingTextConfig{
			ForceQuote:    true,
			DisableColors: false,
			FullTimestamp: true,
		}
	}
}

func validateLogLevel(loggers map[string]LogLevel) {
	for name, lvl := range loggers {
		if !lvl.IsValid() {
			loggers[name] = ERROR
		}
	}
}

func validateHttpPort(cfg *WebServerConfig) {
	if cfg.Base.Port <= 0 || cfg.Base.Port > 65535 {
		cfg.Base.Port = 9000
	}
}

func validateHost(cfg *WebServerConfig) {
	switch {
	case cfg.Base.UseHostname:
		hostname, err := os.Hostname()
		if err != nil || hostname == "" {
			cfg.Base.Host = util.Localhost
		} else {
			cfg.Base.Host = hostname
		}
	case cfg.Base.Host != "":
		// Host già specificato non sono necessarie azioni
	default:
		cfg.Base.Host = util.Localhost
	}
}

func validateBasePath(cfg *WebServerConfig) {
	basePath := strings.TrimSpace(cfg.Base.BasePath)
	if basePath == "" {
		basePath = util.Slash
	}
	if !strings.HasPrefix(basePath, util.Slash) {
		basePath = util.Slash + basePath
	}
	if basePath != util.Slash && strings.HasSuffix(basePath, util.Slash) {
		basePath = strings.TrimSuffix(basePath, util.Slash)
	}
	u := &url.URL{Path: basePath}
	if u.String() != basePath {
		basePath = u.String()
	}
	cfg.Base.BasePath = basePath
}

func validateReadTimeout(cfg *WebServerConfig) {
	if cfg.HTTP.ReadTimeout < 0 {
		cfg.HTTP.ReadTimeout = 10 * time.Second
	}
}

func validateWriteTimeout(cfg *WebServerConfig) {
	if cfg.HTTP.WriteTimeout < 0 {
		cfg.HTTP.WriteTimeout = 10 * time.Second
	}
}

func validateMaxHeaderSizeMB(cfg *WebServerConfig) {
	if cfg.HTTP.MaxHeaderSizeMB <= 0 {
		cfg.HTTP.MaxHeaderSizeMB = 1
	}
}

func validateMaxMultipartMemoryMB(cfg *WebServerConfig) {
	if cfg.HTTP.MaxMultipartMemoryMB <= 0 {
		cfg.HTTP.MaxHeaderSizeMB = 10
	}
}

func validateIdleTimeout(cfg *WebServerConfig) {
	if cfg.HTTP.IdleTimeout < 0 {
		cfg.HTTP.IdleTimeout = 10 * time.Second
	}
}

func validateGracefulShutdownTime(cfg *WebServerConfig) {
	if cfg.HTTP.GracefulShutdownTime <= 0 {
		cfg.HTTP.GracefulShutdownTime = 10 * time.Second
	}
}
