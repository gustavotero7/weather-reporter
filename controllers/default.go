package controllers

import (
	"encoding/json"
	"net/http"
	"weather-reporter/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
	WeatherApp models.IWeatherApp
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) GetWeather() {

	city := c.Ctx.Input.Param(":city")
	country := c.Ctx.Input.Param(":country")

	c.Ctx.ResponseWriter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Ctx.ResponseWriter)
	weatherRespose, err := c.WeatherApp.AskToExternalServiceForWeather(city, country)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(weatherRespose)
}
