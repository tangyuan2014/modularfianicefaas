package server

import (
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

const Prefix  = "/gateway/"
var HostProxy = make(map[string]*httputil.ReverseProxy)

func HandleRequestAndRedirect(writer http.ResponseWriter, request *http.Request) {
	var host string
	if strings.Index(request.URL.Path, Prefix) != 0 {
		log.Println()//TODO
		NotFoundError(writer,request)
		return
	}
	functionName:=request.URL.Path[len(Prefix):]
	if function, ok := ServiceMap[functionName]; ok {
		host = "http://" + function.IpAddress + ":" + function.Port
	} else {
		log.Println()//TODO
		NotFoundError(writer,request)
		return
	}
	log.Println(host)
	if fn, ok := HostProxy[host]; ok {
		fn.ServeHTTP(writer, request)
	} else {
		proxy, err := ServeHttp(host, writer, request)
		if err != nil {
			log.Println()//TODO
			NotFoundError(writer,request)
			return
		}
		HostProxy[host] = proxy
	}
}
