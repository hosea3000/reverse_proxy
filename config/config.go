package config

type Configurations struct {
	Port    int     `mapstructure:"port"`
	Proxies []Proxy `mapstructure:"proxies"`
}

type Proxy struct {
	RouterPath string `mapstructure:"router_path"`
	TargetUrl  string `mapstructure:"target_url"`
}
