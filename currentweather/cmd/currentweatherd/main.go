package main

import (
	"log"
	"net/http"

	"github.com/tangyuan2014/modularfianicefaas/currentweather/cmd/currentweatherd/server"
)

func main() {
	http.HandleFunc("/", server.GetCurrentWeather)
	http.HandleFunc("/health", server.Health)
	http.ListenAndServe(":8080", nil)
	log.Println("Currentweather Service started!")
}
