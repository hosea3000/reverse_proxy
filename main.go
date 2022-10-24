package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	startParams := getStartParam()

	proxy, err := NewProxy(startParams.Host)
	if err != nil {
		fmt.Printf(`代理地址错误`)
		return
	}

	// handle all requests to your server using the proxy
	http.HandleFunc("/", ProxyRequestHandler(proxy))

	fmt.Printf("启动成功: 代理地址=%v 端口号=%v\n", startParams.Host, startParams.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", startParams.Port), nil))
}
