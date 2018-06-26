package controllers

import (
	"beego"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"weather-reporter/models"

	"github.com/stretchr/testify/assert"
)

func MockWeaterAppInfo() *models.WeatherApp {
	weatherApp := models.WeatherApp{
		AppID:    "1508a9a4840a5574c822d70ca2132032",
		Endpoint: "http://api.openweathermap.org/data/2.5/weather?q=%s,co&appid=%s",
	}
	return &weatherApp
}

func TestGetWeather(t *testing.T) {
	r, _ := http.NewRequest("GET", "/weather/Bogota", nil)
	w := httptest.NewRecorder()

	handler := beego.NewControllerRegister()
	handler.Add("/weather/:city", &MainController{WeatherApp: *MockWeaterAppInfo()}, "get:GetWeather")
	handler.ServeHTTP(w, r)

	log.Println("Status code: ", w.Code)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetWeatherWithNoParam(t *testing.T) {
	r, _ := http.NewRequest("GET", "/weather", nil)
	w := httptest.NewRecorder()

	handler := beego.NewControllerRegister()
	handler.Add("/weather/:city", &MainController{WeatherApp: *MockWeaterAppInfo()}, "get:GetWeather")
	handler.ServeHTTP(w, r)

	log.Println("Status code: ", w.Code)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
