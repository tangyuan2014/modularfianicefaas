package main

import (
	"github.com/tangyuan2014/modularfianicefaas/mathoperations/cmd/operationsd/server"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", server.Operation)
	http.HandleFunc("/health", server.Health)
	http.ListenAndServe(":8080", nil)
	log.Println("Mathoperations Service started!")
}
