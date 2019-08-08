package server

import (
	"encoding/json"
	"github.com/sparrc/go-ping"
	"log"
	"net/http"
)

func Ping(writer http.ResponseWriter, request *http.Request) {
	keys, ok := request.URL.Query()["url"]
	if !ok || len(keys[0]) < 1 {
		panic("url param is missing")
	}
	url := keys[0]
	//pinger.SetPrivileged(true)
	pinger, err := ping.NewPinger(url)
	if err != nil {
		log.Println(err.Error())
		notFoundError(writer,request)
		return
	}
	pinger.Count = 3
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	js, err := json.Marshal(stats)
	if err != nil {
		log.Println("")//TODO
		notFoundError(writer,request)
		return
	}
	writer.Write(js)
	return
}
