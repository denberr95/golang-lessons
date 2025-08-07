package config

import (
	"log"
	"main/util"
	"net/url"
	"os"
	"strings"
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
	validateIdleTimeout(cfg)
	validateGracefulShutdownTime(cfg)
}

func validateLoggingConfig(cfg *LoggingConfig) {
	validateLogFormat(cfg.Format)
	validateLogLevel(cfg.Level)
}

func validateLogFormat(format LogFormat) {
	if !format.IsValid() {
		log.Fatal("Unsupported log format value")
	}
	log.Printf("defined log format: %s", format.String())
}

func validateLogLevel(level LogLevel) {
	if !level.IsValid() {
		log.Fatal("Unsupported log level value")
	}
	log.Printf("defined log level: %s", level.String())
}

func validateHttpPort(cfg *WebServerConfig) {
	if cfg.Port <= 0 || cfg.Port > 65535 {
		cfg.Port = 9000
	}
	log.Printf("defined http port: %d", cfg.Port)
}

func validateHost(cfg *WebServerConfig) {
	switch {
	case cfg.UseHostname:
		hostname, err := os.Hostname()
		if err != nil || hostname == "" {
			log.Printf("failed to get system hostname, falling back to localhost: %v", err)
			cfg.Host = util.Localhost
		} else {
			cfg.Host = hostname
		}
	case cfg.Host != "":
		// Host is explicitly set, no action needed
	default:
		cfg.Host = util.Localhost
	}
	log.Printf("defined http host: %s", cfg.Host)
}

func validateBasePath(cfg *WebServerConfig) {
	basePath := strings.TrimSpace(cfg.BasePath)
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
	cfg.BasePath = basePath
	log.Printf("defined base path: %s", cfg.BasePath)
}

func validateReadTimeout(cfg *WebServerConfig) {
	if cfg.ReadTimeout < 0 {
		cfg.ReadTimeout = 10
	}
	log.Printf("defined read timeout: %d seconds", cfg.ReadTimeout)
}

func validateWriteTimeout(cfg *WebServerConfig) {
	if cfg.WriteTimeout < 0 {
		cfg.WriteTimeout = 10
	}
	log.Printf("defined write timeout: %d seconds", cfg.WriteTimeout)
}

func validateMaxHeaderSizeMB(cfg *WebServerConfig) {
	if cfg.MaxHeaderSizeMB <= 0 {
		cfg.MaxHeaderSizeMB = 1
	}
	log.Printf("defined max header size: %d MB", cfg.MaxHeaderSizeMB)
}

func validateIdleTimeout(cfg *WebServerConfig) {
	if cfg.IdleTimeout < 0 {
		cfg.IdleTimeout = 10
	}
	log.Printf("defined idle timeout: %d seconds", cfg.IdleTimeout)
}

func validateGracefulShutdownTime(cfg *WebServerConfig) {
	if cfg.GracefulShutdownTime <= 0 {
		cfg.GracefulShutdownTime = 10
	}
	log.Printf("defined graceful shutdown timeout: %d seconds", cfg.GracefulShutdownTime)
}
