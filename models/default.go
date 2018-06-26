package models

import (
	"fmt"
)

type WeatherApp struct {
	AppID    string
	Endpoint string
}

func (c *WeatherApp) BuildURL(city string, country string) string {
	return fmt.Sprintf(c.Endpoint, country, city, c.AppID)
}

type WeatherResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		Pressure  float64 `json:"pressure"`
		Humidity  float64 `json:"humidity"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		SeaLevel  float64 `json:"sea_level"`
		GrndLevel float64 `json:"grnd_level"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All float64 `json:"all"`
	} `json:"clouds"`
	DT  float64 `json:"dt"`
	Sys struct {
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise float64
		Sunset  float64
	} `json:"sys"`
}
