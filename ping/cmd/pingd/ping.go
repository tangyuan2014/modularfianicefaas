package main

import (
	"encoding/json"
	"fmt"
	"github.com/sparrc/go-ping"
	"net/http"
)

func Ping(writer http.ResponseWriter, request *http.Request) {
	keys, ok := request.URL.Query()["url"]
	if !ok || len(keys[0]) < 1 {
		panic("url param is missing")
	}
	url := keys[0]
	fmt.Println("test1")
	pinger, err := ping.NewPinger(url)
	pinger.SetPrivileged(true)
	if err != nil {
		panic(err)
	}
	fmt.Println("test2")
	pinger.Count = 3
	fmt.Println("test3")
	pinger.Run()                 // blocks until finished
	fmt.Println("test4")
	stats := pinger.Statistics() // get send/receive/rtt stats
	js, err := json.Marshal(stats)
	fmt.Println("test5")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("test4")
	writer.Write(js)
	return
}
func main() {
	http.HandleFunc("/ping/", Ping)
	http.ListenAndServe(":8080", nil)
}
