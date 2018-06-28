package models

import (
	"beego/orm"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
	"weather-reporter/utils"
)

type IWeatherApp interface {
	AskToExternalServiceForWeather(string, string) (WeatherReportRaw, error)
}

type WeatherApp struct {
	AppID    string
	Endpoint string
}

func (c WeatherApp) BuildURL(city string, country string) string {
	return fmt.Sprintf(c.Endpoint, country, city, c.AppID)
}

func (c WeatherApp) AskToExternalServiceForWeather(city string, country string) (WeatherReportRaw, error) {

	url := c.BuildURL(city, country)

	body, err := utils.DoWebRequest(url)

	if err != nil {
		return WeatherReportRaw{}, err
	}

	weatherReportRaw := WeatherReportRaw{}

	err = json.Unmarshal(body, &weatherReportRaw)

	if err != nil {
		return WeatherReportRaw{}, err
	}
	return weatherReportRaw, nil
}

type WeatherReportRaw struct {
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

type WeatherReport struct {
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

func (w *WeatherReport) TableName() string {
	return "reports"
}

func (w *WeatherReport) Parse(raw WeatherReportRaw) {
	w.LocationName = raw.Name + ", " + raw.Sys.Country
	w.Temperature = strconv.FormatFloat((raw.Main.Temp-273.15), 'f', -1, 32) + " °C"
	w.Wind = w.DescribeWindSpeed(raw.Wind.Speed, raw.Wind.Deg)
	w.Pressure = fmt.Sprintf("%s hPa", strconv.FormatFloat(raw.Main.Pressure, 'f', -1, 32))
	w.Humidity = strconv.FormatFloat(raw.Main.Humidity, 'f', -1, 32) + "%"
	w.Sunrise = time.Unix(raw.Sys.Sunrise, 0).Format("15:04")
	w.Sunset = time.Unix(raw.Sys.Sunset, 0).Format("15:04")
	w.GeoCoordinates = "[" + strconv.FormatFloat(raw.Coord.Lat, 'f', -1, 32) + ", " + strconv.FormatFloat(raw.Coord.Lon, 'f', -1, 32) + "]"
	w.RequestedTime = time.Now().UTC() //.Format(time.RFC3339)
}

func (w *WeatherReport) DescribeWindSpeed(speedMS float64, deg float64) string {
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

func (w *WeatherReport) ReadReport() error {
	o := orm.NewOrm()
	o.Using("default")

	qs := o.QueryTable("reports")
	err := qs.Filter("code_name", w.CodeName).One(w)
	if err == nil {
		if time.Since(w.RequestedTime) >= time.Minute*5 { //Expired
			o.Delete(w)
			return errors.New("Expired report")
		}
	}
	return err
}

func (w *WeatherReport) WriteReport() (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	return o.Insert(w)
}

func init() {
	orm.RegisterModel(new(WeatherReport))
}
