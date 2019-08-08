package server

import (
	"io/ioutil"
	"log"
	"net/http"
)

const url = "https://api.openweathermap.org/data/2.5/weather"
const apiKey = "716f03f9d10c7d9c2d3c8f277942940c"

func GetCurrentWeather(writer http.ResponseWriter, request *http.Request) {
	cityName, ok := request.URL.Query()["cityname"]
	if !ok || len(cityName[0]) < 1 {
		log.Println() //TODO
		notFoundError(writer, request)
		return

	}
	curl := url + "?q=" + cityName[0] + "&apikey=" + apiKey
	response, err := http.Get(curl)
	if err != nil {
		log.Println() //TODO
		notFoundError(writer, request)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	writer.Write(body)
}
