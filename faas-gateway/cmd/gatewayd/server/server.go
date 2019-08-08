package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httputil"
)

const Prefix = "/gateway/"

var HostProxy = make(map[string]*httputil.ReverseProxy)

func HandleRequestAndRedirect(writer http.ResponseWriter, request *http.Request) {
	var serviceHost string

	functionPath := request.URL.Path[len(Prefix):]
	if functionService, ok := ServiceMap[functionPath]; ok {
		serviceHost = "http://" + functionService.IPAddress + ":" + functionService.Port
		log.Println("Will redirect " + request.URL.Path + " to " + functionService.Labels + "[" + serviceHost + "]")
	} else {
		logAndWriteError(writer, http.StatusNotFound, errors.New("Failed to find the service with path: "+functionPath))
		return
	}

	if proxy, ok := HostProxy[serviceHost]; ok {
		proxy.ServeHTTP(writer, request)
	} else {
		proxy, err := generateProxy(serviceHost)
		if err != nil {
			logAndWriteError(writer, http.StatusInternalServerError, errors.New("Failed to parse service host: "+serviceHost+" with error: "+err.Error()))
			return
		}
		HostProxy[serviceHost] = proxy
		log.Println("Create new proxy with service host " + serviceHost + " and put into cache")
		proxy.ServeHTTP(writer, request)
	}
}

func DockerStatus(writer http.ResponseWriter, request *http.Request) {
	statusData, err := json.Marshal(ServiceMap)
	if err != nil {
		logAndWriteError(writer, http.StatusInternalServerError, errors.New("Failed to serialize internal service map with error: "+err.Error()))
	}
	writer.Write(statusData)
}

func Health(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("service is up"))
}
