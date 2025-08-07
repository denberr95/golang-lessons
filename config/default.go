package config

import (
	"log"
	"main/util"
	"os"
)

func ApplyDefaults(cfg *Config) {
	applyWebServerDefaults(&cfg.GoApp.WebServer)
	validateLoggingConfig(&cfg.GoApp.Logging)
}

func applyWebServerDefaults(cfg *WebServerConfig) {
	validateHost(cfg)
	validateHttpPort(cfg)
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
