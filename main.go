package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"reverse_proxy/config"
)

func main() {
	//startParams := getStartParam()

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	var configuration config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	for _, p := range configuration.Proxies {
		proxy, err := NewProxy(p.TargetUrl)
		if err != nil {
			fmt.Printf(`代理地址错误`)
			return
		}

		http.HandleFunc(p.RouterPath, ProxyRequestHandler(proxy))
	}

	fmt.Printf("启动成功: 端口号=%v\n", configuration.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port), nil))
}
