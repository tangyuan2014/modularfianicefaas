package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/tangyuan2014/modularfianicefaas/faas-gateway/cmd/gatewayd/server"
	"net/http"
)

func pollingStatusOfService(){
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(server.GetContainerStatus)
	<-s.Start()
}

func main() {
	go pollingStatusOfService()
	http.HandleFunc(server.Prefix, server.HandleRequestAndRedirect)
	err:=http.ListenAndServe(":80", nil)
	if err!=nil{
		panic("")//TODO
	}

}
