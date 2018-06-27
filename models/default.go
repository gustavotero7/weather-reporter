package models

import (
	"beego/orm"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"weather-reporter/utils"
)

type IWeatherApp interface {
	AskToExternalServiceForWeather(string, string) (WeatherResponse, error)
}

type WeatherApp struct {
	AppID    string
	Endpoint string
}

func (c WeatherApp) BuildURL(city string, country string) string {
	return fmt.Sprintf(c.Endpoint, country, city, c.AppID)
}

func (c WeatherApp) AskToExternalServiceForWeather(city string, country string) (WeatherResponse, error) {

	url := c.BuildURL(city, country)

	body, err := utils.DoWebRequest(url)

	if err != nil {
		return WeatherResponse{}, err
	}

	weatherResponse := WeatherResponse{}

	err = json.Unmarshal(body, &weatherResponse)

	if err != nil {
		return WeatherResponse{}, err
	}
	return weatherResponse, nil
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
		Sunrise int64
		Sunset  int64
	} `json:"sys"`
}

type WeatherParsed struct {
	ID             int       `orm:"column(id)"`
	CodeName       string    `orm:"column(code_name)"`
	LocationName   string    `orm:"column(location_name)"`
	Temperature    string    `orm:"column(temperature)"`
	Wind           string    `orm:"column(wind)"`
	Pressure       string    `orm:"column(pressure)"`
	Humidity       string    `orm:"column(humidity)"`
	Sunrise        string    `orm:"column(sunrise)"`
	Sunset         string    `orm:"column(sunset)"`
	GeoCoordinates string    `orm:"column(geo_coordinates)"`
	RequestedTime  time.Time `orm:"column(requested_time)"`
}

func (w *WeatherParsed) TableName() string {
	return "reports"
}

func (w *WeatherParsed) Parse(wr WeatherResponse) {
	w.LocationName = wr.Name + ", " + wr.Sys.Country
	w.Temperature = strconv.FormatFloat((wr.Main.Temp-273.15), 'f', -1, 32) + " °C"
	w.Wind = w.DescribeWindSpeed(wr.Wind.Speed, wr.Wind.Deg)
	w.Pressure = fmt.Sprintf("%s hPa", strconv.FormatFloat(wr.Main.Pressure, 'f', -1, 32))
	w.Humidity = strconv.FormatFloat(wr.Main.Humidity, 'f', -1, 32) + "%"
	w.Sunrise = time.Unix(wr.Sys.Sunrise, 0).Format("15:04")
	w.Sunset = time.Unix(wr.Sys.Sunset, 0).Format("15:04")
	w.GeoCoordinates = "[" + strconv.FormatFloat(wr.Coord.Lat, 'f', -1, 32) + ", " + strconv.FormatFloat(wr.Coord.Lon, 'f', -1, 32) + "]"
	w.RequestedTime = time.Now() //.Format(time.RFC3339)
}

func (w *WeatherParsed) DescribeWindSpeed(speedMS float64, deg float64) string {
	descriptions := []string{
		"Calm",
		"Light Air",
		"Ligth breeze",
		"Gentle breeze",
		"Moderate breeze",
		"Fresh breeze",
		"Strong breeze",
		"High wind",
		"Gale, Fresh gale",
		"Strong gale",
		"Storm, whole gale",
		"Violent storm",
		"Huricane",
	}

	directions := []string{
		"North",
		"Est",
		"South",
		"West",
	}

	var windSpeedDescription string
	var windDegDescription string
	speed := (speedMS * 18) / 5
	if speed < 1 {
		windSpeedDescription = descriptions[0]
	} else if speed >= 1 && speed < 6 {
		windSpeedDescription = descriptions[1]
	} else if speed >= 6 && speed < 12 {
		windSpeedDescription = descriptions[2]
	} else if speed >= 12 && speed < 20 {
		windSpeedDescription = descriptions[3]
	} else if speed >= 20 && speed < 29 {
		windSpeedDescription = descriptions[4]
	} else if speed >= 29 && speed < 39 {
		windSpeedDescription = descriptions[5]
	} else if speed >= 39 && speed < 50 {
		windSpeedDescription = descriptions[6]
	} else if speed >= 50 && speed < 62 {
		windSpeedDescription = descriptions[7]
	} else if speed >= 62 && speed < 75 {
		windSpeedDescription = descriptions[8]
	} else if speed >= 75 && speed < 89 {
		windSpeedDescription = descriptions[9]
	} else if speed >= 89 && speed < 103 {
		windSpeedDescription = descriptions[10]
	} else if speed >= 103 && speed < 118 {
		windSpeedDescription = descriptions[11]
	} else if speed >= 118 {
		windSpeedDescription = descriptions[12]
	}

	if deg > 315 || deg < 45 {
		windDegDescription = directions[0] + "-" + directions[2]
	} else if deg >= 45 && deg < 135 {
		windDegDescription = directions[1] + "-" + directions[3]
	} else if deg >= 135 && deg < 225 {
		windDegDescription = directions[2] + "-" + directions[0]
	} else if deg >= 225 && deg < 315 {
		windDegDescription = directions[3] + "-" + directions[1]
	}
	windDegDescription += ", " + strconv.FormatFloat(deg, 'f', -1, 32) + "°"

	return windSpeedDescription + ", " + strconv.FormatFloat(speedMS, 'f', -1, 32) + " m/s (" + strconv.FormatFloat(speed, 'f', -1, 32) + " Km/h), " + windDegDescription
}

func init() {
	orm.RegisterModel(new(WeatherParsed))
}
