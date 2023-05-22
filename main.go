package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
	"reverse_proxy/config"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Printf(`初始化配置错误: %+v \n`, err)
		return
	}

	// 读取证书和私钥文件
	certFile := config.Configuration.SSL.CertFile
	keyFile := config.Configuration.SSL.KeyFile

	// 创建TLS配置
	tlsConfig, err := createTLSConfig(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	// 创建HTTPS服务器
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", config.Configuration.Port),
		TLSConfig: tlsConfig,
	}

	// 处理请求
	http.HandleFunc("/", HandleRequestAndRedirect)

	// 启动服务器
	log.Fatal(server.ListenAndServeTLS(certFile, keyFile))
}

// 创建TLS配置
func createTLSConfig(certFile, keyFile string) (*tls.Config, error) {
	// 创建一个CA池，并加载客户端证书
	caCert, err := os.ReadFile("./ssl/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// 读取证书和私钥文件
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	// 创建TLS配置
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caCertPool,
	}

	return tlsConfig, nil
}
