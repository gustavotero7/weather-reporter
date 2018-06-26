package routers

import (
	"weather-reporter/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/weather/:city", &controllers.MainController{}, "get:GetWeather")
}
