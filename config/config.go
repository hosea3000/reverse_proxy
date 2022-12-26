package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Configurations struct {
	Port    int     `mapstructure:"port"`
	Https   bool    `mapstructure:"https"`
	Proxies []Proxy `mapstructure:"proxies"`
}

type Proxy struct {
	RouterPath string `mapstructure:"router_path"`
	TargetUrl  string `mapstructure:"target_url"`
}

var Configuration Configurations

func Init() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return errors.WithStack(err)
	}
	err := viper.Unmarshal(&Configuration)

	return errors.WithStack(err)
}
