package routers

import (
	"weather-reporter/controllers"
	"weather-reporter/models"

	"github.com/astaxie/beego"
)

func init() {

	weatherApp := models.WeatherApp{
		AppID:    beego.AppConfig.String("appid"),
		Endpoint: beego.AppConfig.String("endpoint"),
	}

	beego.Router("/", &controllers.MainController{})
	beego.Router("/weather/:city/:country", &controllers.MainController{WeatherApp: weatherApp}, "get:GetWeather")
}
