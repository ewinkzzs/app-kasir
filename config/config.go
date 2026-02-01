package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func LoadConfig() (Config, error) {
	var config Config

	// baca environment variable dari sistem
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// cek apakah file .env ada
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		// baca isi file .env
		if err := viper.ReadInConfig(); err == nil {
			log.Println("Using .env file")
		} else {
			log.Println("Failed to read .env file, fallback to system env")
		}
	} else {
		log.Println("No .env file found, fallback to system env")
	}

	// mapping ke struct Config
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
