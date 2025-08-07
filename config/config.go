package config

import (
	"flag"
	"log"
	"main/util"
	"strings"

	"github.com/spf13/viper"
)

var cfg Config

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

func Init(flags *ProgramFlags) {

	viper.SetConfigName(flags.ConfigFileName)
	viper.SetConfigType(flags.FileType)
	viper.AddConfigPath(flags.FilePath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(util.Dot, util.Underscore))
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to decode into struct: %v", err)
	}

	applyConfiguration(&cfg)
}

func GetConfig() *Config {
	return &cfg
}
