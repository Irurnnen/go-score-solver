package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var config *viper.Viper

func init() {
	config = viper.New()
	config.SetConfigName("score_config")
	config.SetConfigType("yaml")
	config.AddConfigPath("/run/secrets")

	err := config.ReadInConfig()
	if err != nil {
		log.WithError(err).Fatal("Failed to load configuration")
	}

	config.SetDefault("port", "6874")
	config.SetDefault("host", "0.0.0.0")
	config.SetDefault("release", "release")
}

func GetServer() string {
	return config.GetString("host") + ":" + config.GetString("port")
}

func GetRelease() string {
	return config.GetString("release")
}
