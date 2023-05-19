package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Configurations struct {
	Port    int     `mapstructure:"port"`
	Proxies []Proxy `mapstructure:"proxies"`
	SSL     SSL     `mapstructure:"ssl"`
}

type Proxy struct {
	RouterPath string `mapstructure:"router_path"`
	TargetUrl  string `mapstructure:"target_url"`
}

type SSL struct {
	CertFile string `mapstructure:"cert_file"`
	KeyFile  string `mapstructure:"key_file"`
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
