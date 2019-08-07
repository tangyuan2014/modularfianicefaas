package models

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)


func ServeHttp(target string, writer http.ResponseWriter, request *http.Request) *httputil.ReverseProxy {
	targetUrl, err := url.Parse(target)
	if err!=nil{
		log.Println("url fail")
		return nil
	}
	proxy:=httputil.NewSingleHostReverseProxy(targetUrl)
	request.URL.Host = targetUrl.Host
	request.URL.Scheme = targetUrl.Scheme
	request.Header.Set("X-Forwarded-Host", request.Header.Get("Host"))
	request.Host = targetUrl.Host
	proxy.ServeHTTP(writer,request)
	return proxy
}

