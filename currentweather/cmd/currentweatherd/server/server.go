package server

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "https://api.openweathermap.org/data/2.5/weather"
const apiKey = "716f03f9d10c7d9c2d3c8f277942940c"
const cityName = "cityname"

func GetCurrentWeather(writer http.ResponseWriter, request *http.Request) {
	cityName, err := validateAndGetInput(request, cityName)
	if err != nil {
		logAndWriteError(writer, http.StatusBadRequest, err)
		log.Println("Input validation failed with error: " + err.Error())
		return
	}

	curl := url + "?q=" + cityName + "&apikey=" + apiKey
	response, err := http.Get(curl)
	if err != nil {
		logAndWriteError(writer, http.StatusInternalServerError, errors.New("Failed to get response from Openweather with error:"+err.Error()))
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	writer.Write(body)
}

func validateAndGetInput(request *http.Request, paramKey string) (string, error) {
	param, ok := request.URL.Query()[paramKey]
	if !ok {
		return "", errors.New("Failed to get param value from key: " + paramKey)
	} else if len(param) != 1 {
		return "", errors.New("Please provide one and only one city")
	}

	return param[0], nil
}

func Health(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("service is up"))
}
