package server

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCurrentWeather(t *testing.T) {
	req, err := http.NewRequest("GET", "/?cityname=London", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	resp := httptest.NewRecorder()
	GetCurrentWeather(resp, req)
	expectresp := `{"coord":{"lon":-0.13,"lat":51.51},"weather":[{"id":521,"main":"Rain","description":"shower rain","icon":"09d"},{"id":311,"main":"Drizzle","description":"drizzle rain","icon":"09d"},{"id":721,"main":"Haze","description":"haze","icon":"50d"}],"base":"stations","main":{"temp":291.65,"pressure":1002,"humidity":93,"temp_min":290.93,"temp_max":292.59},"visibility":10000,"wind":{"speed":4.6,"deg":160},"rain":{"1h":0.51},"clouds":{"all":40},"dt":1565330127,"sys":{"type":1,"id":1412,"message":0.01,"country":"GB","sunrise":1565325358,"sunset":1565379376},"timezone":3600,"id":2643743,"name":"London","cod":200}`
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, expectresp, resp.Body.String())
}

func TestGetCurrentWeather2(t *testing.T) {
	req, err := http.NewRequest("GET", "/?ityname=London", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	resp := httptest.NewRecorder()
	GetCurrentWeather(resp, req)
	expectresp :=`{"Code":400,"Message":"Failed to get param value from key: cityname"}`
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Equal(t, expectresp, resp.Body.String())
}
