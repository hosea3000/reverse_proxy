package main

import (
	"fmt"
	"log"
	"net/http"
	"reverse_proxy/config"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Printf(`初始化配置错误: %+v \n`, err)
		return
	}

	http.HandleFunc("/", HandleRequestAndRedirect)

	port := config.Configuration.Port
	fmt.Printf("启动成功: 端口号=%v\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
