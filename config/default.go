package config

func ApplyDefaults(cfg *Config) {
	applyWebServerDefaults(&cfg.GoApp.WebServer)
	validateLoggingConfig(&cfg.GoApp.Logging)
}

func applyWebServerDefaults(cfg *WebServerConfig) {
	if cfg.Host == "" {
		cfg.Host = "localhost"
	}
	if cfg.Port == 0 {
		cfg.Port = 9000
	}
}

func validateLoggingConfig(cfg *LoggingConfig) {
	if !cfg.Format.IsValid() {
		panic("unsupported log format")
	}
	if !cfg.Level.IsValid() {
		panic("Unsupported log level")
	}
}
