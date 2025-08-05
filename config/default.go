package config

func ApplyDefaults(cfg *Config) {
	applyWebServerDefaults(&cfg.GoApp.WebServer)
	applyLoggingDefaults(&cfg.GoApp.Logging)
}

func applyWebServerDefaults(ws *WebServerConfig) {
	if ws.Host == "" {
		ws.Host = "0.0.0.0"
	}
	if ws.Port == 0 {
		ws.Port = 9000
	}
}

func applyLoggingDefaults(log *LoggingConfig) {
	if log.Format == "" {
		log.Format = "text"
	}
	if log.Level == "" {
		log.Level = "info"
	}
}