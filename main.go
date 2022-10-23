package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// NewProxy takes target host and creates a reverse proxy
func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		modifyRequest(req)
	}

	//proxy.ModifyResponse = modifyResponse()
	//proxy.ErrorHandler = errorHandler()
	return proxy, nil
}

func modifyRequest(req *http.Request) {
	req.Header.Set("X-Proxy", "Simple-Reverse-Proxy")
}

func errorHandler() func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, req *http.Request, err error) {
		fmt.Printf("Got error while modifying response: %v \n", err)
		return
	}
}

func modifyResponse() func(*http.Response) error {
	return func(resp *http.Response) error {
		return errors.New("response body is invalid")
	}
}

// ProxyRequestHandler handles the http request using proxy
func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

type StartParams struct {
	Host string
	Port int
}

func getStartParam() StartParams {
	// 主机名
	var host string
	// 端口号
	var port int

	flag.StringVar(&host, "host", "http://127.0.0.1", "主机名,默认 http://127.0.0.1")
	flag.IntVar(&port, "port", 8080, "端口号,默认为8080")

	// 从arguments中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()

	return StartParams{
		Host: host,
		Port: port,
	}
}

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
