package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"weather-reporter/models"
	"weather-reporter/utils"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
	WeatherApp models.WeatherApp
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) GetWeather() {

	city := c.Ctx.Input.Param(":city")
	country := c.Ctx.Input.Param(":country")

	url := c.WeatherApp.BuildURL(city, country)

	body, err := utils.DoWebRequest(url)

	if err != nil {
		log.Fatal(err)
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	weatherResponse := models.WeatherResponse{}

	err = json.Unmarshal(body, &weatherResponse)

	if err != nil {
		log.Fatal(err)
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Header().Add("Content-Type", "application/json")
	//c.Ctx.ResponseWriter.Write(body)
	encoder := json.NewEncoder(c.Ctx.ResponseWriter)
	err = encoder.Encode(weatherResponse)

	if err != nil {
		log.Fatal(err)
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
}
