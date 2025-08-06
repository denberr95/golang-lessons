package config

import (
	"flag"
	"fmt"
	"main/util"
	"strings"

	"github.com/spf13/viper"
)

func ParseFlags() *ProgramFlags {
	fileName := flag.String("config-file-name", "config.default", "Configuration file name")
	fileType := flag.String("config-file-type", "yaml", "Type of the configuration file")
	filePath := flag.String("config-file-path", ".", "Path to the configuration file directory")
	flag.Parse()

	return &ProgramFlags{
		ConfigFileName: *fileName,
		FileType:       *fileType,
		FilePath:       *filePath,
	}
}

func LoadConfig(flags *ProgramFlags) (*Config, error) {
	var cfg Config

	viper.SetConfigName(flags.ConfigFileName)
	viper.SetConfigType(flags.FileType)
	viper.AddConfigPath(flags.FilePath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(util.Dot, util.Underscore))
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	ApplyDefaults(&cfg)

	return &cfg, nil
}
