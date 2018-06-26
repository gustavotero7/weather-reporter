package controllers

import (
	"beego"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"weather-reporter/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func NewWeaterAppMock() *MockWeaterApp {
	mockWeatherApp := new(MockWeaterApp)
	return mockWeatherApp
}

type MockWeaterApp struct {
	mock.Mock
	AppID    string
	Endpoint string
}

func (m MockWeaterApp) AskToExternalServiceForWeather(city string, country string) (models.WeatherResponse, error) {
	args := m.Called(city, country)
	return args.Get(0).(models.WeatherResponse), args.Error(1)
}

func TestGetWeather(t *testing.T) {

	mockWeatherApp := NewWeaterAppMock()
	mockWeatherApp.On("AskToExternalServiceForWeather", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(models.WeatherResponse{}, nil)

	r, _ := http.NewRequest("GET", "/weather/unknownCountry/unknownCity", nil)
	w := httptest.NewRecorder()
	handler := beego.NewControllerRegister()
	handler.Add("/weather/:country/:city", &MainController{WeatherApp: mockWeatherApp}, "get:GetWeather")
	handler.ServeHTTP(w, r)

	mockWeatherApp.AssertExpectations(t)
}

func TestGetWeatherWithNoParam(t *testing.T) {

	mockWeatherApp := NewWeaterAppMock()
	mockWeatherApp.On("AskToExternalServiceForWeather", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(models.WeatherResponse{}, errors.New("Server error"))

	r, _ := http.NewRequest("GET", "/weather/unknownCountry/unknownCity", nil)
	w := httptest.NewRecorder()

	handler := beego.NewControllerRegister()
	handler.Add("/weather/:country/:city", &MainController{WeatherApp: mockWeatherApp}, "get:GetWeather")
	handler.ServeHTTP(w, r)

	log.Println("Status code: ", w.Code)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
