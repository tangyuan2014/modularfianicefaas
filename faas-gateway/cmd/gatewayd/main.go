package main

import (
	"log"
	"net/http"

	"github.com/jasonlvhit/gocron"
	"github.com/tangyuan2014/modularfianicefaas/faas-gateway/cmd/gatewayd/server"
)

func pollServiceStatus() {
	s := gocron.NewScheduler()
	s.Every(5).Seconds().Do(server.GetContainerStatus)
	<-s.Start()
}

func main() {
	go pollServiceStatus()
	http.HandleFunc(server.Prefix, server.HandleRequestAndRedirect)
	http.HandleFunc("/status/", server.DockerStatus)
	http.HandleFunc("/health/", server.Health)
	http.ListenAndServe(":80", nil)
	log.Println("Gateway Service started!")
}
