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

	// baca environment variable
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// cek file .env hanya di lokal
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		if err := viper.ReadInConfig(); err == nil {
			log.Println("Using .env file")
		}
	}

	// isi struct Config dari env
	config.DBConn = strings.TrimSpace(viper.GetString("DB_CONN"))
	config.Port = strings.TrimSpace(viper.GetString("PORT"))

	return config, nil
}
