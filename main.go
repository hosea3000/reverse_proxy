package main

import (
	"crypto/tls"
	"flag"
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

	port := config.Configuration.Port
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("启动成功: 端口号=%v\n", port)

	if config.Configuration.Https {
		RunHttps(addr)
	} else {
		RunHttp(addr)
	}
}

func RunHttp(addr string) {
	http.HandleFunc("/", HandleRequestAndRedirect)

	log.Fatal(http.ListenAndServe(addr, nil))
}

func RunHttps(addr string) {
	certFile := "./cert/cert.pem"
	keyFile := "./cert/key.pem"
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleRequestAndRedirect)

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS13,
			PreferServerCipherSuites: true,
		},
	}

	err := srv.ListenAndServeTLS(certFile, keyFile)
	log.Fatal(err)
}
