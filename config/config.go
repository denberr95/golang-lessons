package config

import (
	"fmt"	
	"strings"

	"github.com/spf13/viper"
)

type WebServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type LoggingConfig struct {
	Format string `mapstructure:"format"`
	Level  string `mapstructure:"level"`
}

type Config struct {
	GoApp struct {
		WebServer WebServerConfig `mapstructure:"webserver"`
		Logging   LoggingConfig   `mapstructure:"logging"`
	} `mapstructure:"goapp"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	
	viper.SetConfigName("config.local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	ApplyDefaults(&cfg)

	return &cfg, nil
}