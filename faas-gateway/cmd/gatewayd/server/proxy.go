package server

import (
	"log"
	"net/http/httputil"
	"net/url"
)

func generateProxy(target string) (*httputil.ReverseProxy, error) {
	log.Println(target)
	targetUrl, err := url.Parse(target)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(targetUrl)
	return proxy, nil
}
