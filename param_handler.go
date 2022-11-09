package main

import "flag"

type StartParams struct {
	Host string
	Port int
}

func getStartParam() StartParams {
	// 主机名
	var host string
	// 端口号
	var port int

	flag.StringVar(&host, "host", "http://127.0.0.1", "反向代理的地址")
	flag.IntVar(&port, "port", 8080, "服务启动端口号")

	// 从arguments中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()

	return StartParams{
		Host: host,
		Port: port,
	}
}
