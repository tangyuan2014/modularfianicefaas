package main

import (
	"net/http"

	"github.com/tangyuan2014/modularfianicefaas/ping/cmd/pingd/server"
)

func main() {
	http.HandleFunc("/", server.Ping)
	http.ListenAndServe(":8080", nil)
}
