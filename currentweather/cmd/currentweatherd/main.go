package main

import (
	"net/http"

	"github.com/tangyuan2014/modularfianicefaas/currentweather/cmd/currentweatherd/server"
)

func main() {
	http.HandleFunc("/", server.GetCurrentWeather)
	http.ListenAndServe(":8080", nil)
}
