package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	AppPort     string
	Debug       bool
	AutoMigrate bool
	DbURL       string
}

var Cfg Config

func InitConfig() {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No config file found, loading from ENV...")
	}

	Cfg = Config{
		AppPort:     viper.GetString("APP_PORT"),
		Debug:       viper.GetBool("DEBUG"),
		AutoMigrate: viper.GetBool("AUTO_MIGRATE"),
		DbURL:       viper.GetString("DB_URL"),
	}
}
