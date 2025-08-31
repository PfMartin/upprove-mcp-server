package config

import (
	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

type Config struct {
	DBName     string `mapstructure:"UPPROVE_DB"`
	DBUser     string `mapstructure:"UPPROVE_USER"`
	DBPassword string `mapstructure:"UPPROVE_PWD"`
	DBURI      string `mapstructure:"UPPROVE_URI"`
}

func NewConfig(configPath string, configName string) (config Config, err error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		log.Fatal().Msgf("failed to read config: %s", err)
		return
	}

	err = viper.Unmarshal(&config)

	return config, err
}
