package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"Port"`
}

var env = []string{
	"Port",
}

func LoadConfig() (Config, error) {
	var config Config
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, v := range env {
		if err := viper.BindEnv(v); err != nil {
			println(v)
			return config, err
		}
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err

	}
	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}
	return config, nil
}
