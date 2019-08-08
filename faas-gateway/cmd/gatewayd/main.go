package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/tangyuan2014/modularfianicefaas/faas-gateway/cmd/gatewayd/servicehandling"
	"net/http"
)

func pollingStatusOfService(){
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(servicehandling.GetContainerStatus)
	<-s.Start()
}

func main() {
	go pollingStatusOfService()
	http.HandleFunc(servicehandling.Prefix, servicehandling.HandleRequestAndRedirect)
	err:=http.ListenAndServe(":80", nil)
	if err!=nil{
		panic("")//TODO
	}

}