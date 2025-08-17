package config

import (
	"log"
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
	validateLogLevel(cfg.Base.Level)
}

func validateLogFormat(cfg *LoggingConfig) {
	if cfg.Text == nil && cfg.JSON == nil {
		cfg.Text = &LoggingTextConfig{
			ForceQuote:    true,
			DisableColors: false,
			FullTimestamp: true,
		}
		log.Println("Formato Log definito: text")
	}
}

func validateLogLevel(level LogLevel) {
	if !level.IsValid() {
		log.Fatal("Log level configurato non supportato")
	}
	log.Printf("Log level definito: %s", level.String())
}

func validateHttpPort(cfg *WebServerConfig) {
	if cfg.Base.Port <= 0 || cfg.Base.Port > 65535 {
		cfg.Base.Port = 9000
	}
	log.Printf("Porta HTTP definita: %d", cfg.Base.Port)
}

func validateHost(cfg *WebServerConfig) {
	switch {
	case cfg.Base.UseHostname:
		hostname, err := os.Hostname()
		if err != nil || hostname == "" {
			log.Printf("Errore nel recupero dell'hostname di sistema, utilizzo 'localhost': %v", err)
			cfg.Base.Host = util.Localhost
		} else {
			cfg.Base.Host = hostname
		}
	case cfg.Base.Host != "":
		// Host is explicitly set, no action needed
	default:
		cfg.Base.Host = util.Localhost
	}
	log.Printf("HTTP Host definito: %s", cfg.Base.Host)
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
	log.Printf("HTTP base path definito: %s", cfg.Base.BasePath)
}

func validateReadTimeout(cfg *WebServerConfig) {
	if cfg.HTTP.ReadTimeout < 0 {
		cfg.HTTP.ReadTimeout = 10 * time.Second
	}
	log.Printf("Definito HTTP read timeout: %s", cfg.HTTP.ReadTimeout)
}

func validateWriteTimeout(cfg *WebServerConfig) {
	if cfg.HTTP.WriteTimeout < 0 {
		cfg.HTTP.WriteTimeout = 10 * time.Second
	}
	log.Printf("Definito HTTP write timeout: %s", cfg.HTTP.WriteTimeout)
}

func validateMaxHeaderSizeMB(cfg *WebServerConfig) {
	if cfg.HTTP.MaxHeaderSizeMB <= 0 {
		cfg.HTTP.MaxHeaderSizeMB = 1
	}
	log.Printf("Definita HTTP max header size: %d MB", cfg.HTTP.MaxHeaderSizeMB)
}

func validateMaxMultipartMemoryMB(cfg *WebServerConfig) {
	if cfg.HTTP.MaxMultipartMemoryMB <= 0 {
		cfg.HTTP.MaxHeaderSizeMB = 10
	}
	log.Printf("Definita HTTP max multipart memory size: %d MB", cfg.HTTP.MaxMultipartMemoryMB)
}

func validateIdleTimeout(cfg *WebServerConfig) {
	if cfg.HTTP.IdleTimeout < 0 {
		cfg.HTTP.IdleTimeout = 10 * time.Second
	}
	log.Printf("Definito HTTP idle timeout: %s", cfg.HTTP.IdleTimeout)
}

func validateGracefulShutdownTime(cfg *WebServerConfig) {
	if cfg.HTTP.GracefulShutdownTime <= 0 {
		cfg.HTTP.GracefulShutdownTime = 10 * time.Second
	}
	log.Printf("Definito HTTP graceful shutdown timeout: %s", cfg.HTTP.GracefulShutdownTime)
}
