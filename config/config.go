package config

import (
	"flag"
	"log"
	"main/util"
	"strings"

	"github.com/spf13/viper"
)

var cfg Config
var flags *ProgramFlags = parseFlags()

func parseFlags() *ProgramFlags {
	fileName := flag.String("config-file-name", "config.default", "Nome file di configurazione")
	fileType := flag.String("config-file-type", "yaml", "Tipo di file di configurazione")
	filePath := flag.String("config-file-path", ".", "Percorso file di configurazione")
	flag.Parse()

	return &ProgramFlags{
		ConfigFileName: *fileName,
		FileType:       *fileType,
		FilePath:       *filePath,
	}
}

func Load() {

	viper.SetConfigName(flags.ConfigFileName)
	viper.SetConfigType(flags.FileType)
	viper.AddConfigPath(flags.FilePath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Errore lettura file di configurazione: %v", err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(util.Dot, util.Underscore))
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Impossibile decodificare file di configurazione: %v", err)
	}

	applyConfiguration(&cfg)
}

func GetConfig() *Config {
	return &cfg
}
