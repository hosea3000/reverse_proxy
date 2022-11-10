package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"reverse_proxy/config"
	"strings"
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

	proxy.ModifyResponse = modifyResponse()
	return proxy, nil
}

func modifyRequest(req *http.Request) {
	realIP := GetRealIP(req)
	req.Header.Set("X-Proxy", "Simple-Reverse-Proxy")
	req.Header.Set("X-Client-Real-IP", realIP)
}

func modifyResponse() func(*http.Response) error {
	return func(resp *http.Response) error {
		return nil
	}
}

func GetRealIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-IP")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarder-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func getProxyUrl(req *http.Request) string {
	uri := req.RequestURI
	proxyCondition := strings.ToUpper(uri)

	defaultUrl := ""
	for _, p := range config.Configuration.Proxies {
		if p.RouterPath == "" {
			defaultUrl = p.RouterPath
			continue
		}
		if strings.HasPrefix(proxyCondition, strings.ToUpper(p.RouterPath)) {
			return p.TargetUrl
		}
	}

	return defaultUrl
}

func HandleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	targetUrl := getProxyUrl(req)

	proxy, _ := NewProxy(targetUrl)
	proxy.ServeHTTP(res, req)
}
