package server

import (
	"log"
	"net/http/httputil"
	"net/url"
)

func generateProxy(target string) (*httputil.ReverseProxy, error) {
	log.Println(target)
	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	return proxy, nil
}
