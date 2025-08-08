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
	validateMaxMultipartMemoryMB(cfg)
	validateIdleTimeout(cfg)
	validateGracefulShutdownTime(cfg)
}

func validateLoggingConfig(cfg *LoggingConfig) {
	validateLogFormat(cfg.Format)
	validateLogLevel(cfg.Level)
}

func validateLogFormat(format LogFormat) {
	if !format.IsValid() {
		log.Fatal("Log format configurato non supportato")
	}
	log.Printf("Log format definito: %s", format.String())
}

func validateLogLevel(level LogLevel) {
	if !level.IsValid() {
		log.Fatal("Log level configurato non supportato")
	}
	log.Printf("Log level definito: %s", level.String())
}

func validateHttpPort(cfg *WebServerConfig) {
	if cfg.Port <= 0 || cfg.Port > 65535 {
		cfg.Port = 9000
	}
	log.Printf("Porta HTTP definita: %d", cfg.Port)
}

func validateHost(cfg *WebServerConfig) {
	switch {
	case cfg.UseHostname:
		hostname, err := os.Hostname()
		if err != nil || hostname == "" {
			log.Printf("Errore nel recupero dell'hostname di sistema, utilizzo 'localhost': %v", err)
			cfg.Host = util.Localhost
		} else {
			cfg.Host = hostname
		}
	case cfg.Host != "":
		// Host is explicitly set, no action needed
	default:
		cfg.Host = util.Localhost
	}
	log.Printf("HTTP Host definito: %s", cfg.Host)
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
	log.Printf("HTTP base path definito: %s", cfg.BasePath)
}

func validateReadTimeout(cfg *WebServerConfig) {
	if cfg.ReadTimeout < 0 {
		cfg.ReadTimeout = 10
	}
	log.Printf("Definito HTTP read timeout: %d secondi", cfg.ReadTimeout)
}

func validateWriteTimeout(cfg *WebServerConfig) {
	if cfg.WriteTimeout < 0 {
		cfg.WriteTimeout = 10
	}
	log.Printf("Definito HTTP write timeout: %d secondi", cfg.WriteTimeout)
}

func validateMaxHeaderSizeMB(cfg *WebServerConfig) {
	if cfg.MaxHeaderSizeMB <= 0 {
		cfg.MaxHeaderSizeMB = 1
	}
	log.Printf("Definita HTTP max header size: %d MB", cfg.MaxHeaderSizeMB)
}

func validateMaxMultipartMemoryMB(cfg *WebServerConfig) {
	if cfg.MaxMultipartMemoryMB <= 0 {
		cfg.MaxHeaderSizeMB = 10
	}
	log.Printf("Definita HTTP max multipart memory size: %d MB", cfg.MaxMultipartMemoryMB)
}

func validateIdleTimeout(cfg *WebServerConfig) {
	if cfg.IdleTimeout < 0 {
		cfg.IdleTimeout = 10
	}
	log.Printf("Definito HTTP idle timeout: %d secondi", cfg.IdleTimeout)
}

func validateGracefulShutdownTime(cfg *WebServerConfig) {
	if cfg.GracefulShutdownTime <= 0 {
		cfg.GracefulShutdownTime = 10
	}
	log.Printf("Definito HTTP graceful shutdown timeout: %d second", cfg.GracefulShutdownTime)
}
