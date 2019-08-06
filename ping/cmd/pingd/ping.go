package main

import (
	"encoding/json"
	"github.com/sparrc/go-ping"
	"net/http"
)

func Ping(writer http.ResponseWriter, request *http.Request) {
	keys, ok := request.URL.Query()["url"]
	if !ok || len(keys[0]) < 1 {
		panic("url param is missing")
	}
	url := 	//pinger.SetPrivileged(true)
		keys[0]
	pinger, err := ping.NewPinger(url)
	if err != nil {
		panic(err)
	}
 	pinger.Count = 3
 	pinger.Run()                 // blocks until finished
 	stats := pinger.Statistics() // get send/receive/rtt stats
	js, err := json.Marshal(stats)

	if err != nil {
		writer.Write([]byte(err.Error()))
		return
	}
 	writer.Write(js)
	return
}
func main() {
	http.HandleFunc("/ping/", Ping)
	http.ListenAndServe(":8080", nil)
}
