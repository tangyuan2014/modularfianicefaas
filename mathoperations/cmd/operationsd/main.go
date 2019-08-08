package main

import (
	"github.com/tangyuan2014/modularfianicefaas/mathoperations/cmd/operationsd/server"
	"net/http"
)

func main() {
	http.HandleFunc("/", server.Operation)
	http.ListenAndServe(":8080", nil)
}
