package main

import (
	"github.com/tangyuan2014/modularfianicefaas/ping/cmd/pingd/server"
	"net/http"
)

func main() {
	http.HandleFunc("/", server.Ping)
	http.ListenAndServe(":8080", nil)
}
