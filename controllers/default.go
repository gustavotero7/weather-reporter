package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
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

	// Create empty report object
	weatherReport := &models.WeatherReport{
		CodeName: strings.ToLower(city) + strings.ToLower(country),
	}

	// Ask to database for a saved report
	// If got error (due record not exist or it's expired)
	// then ask external service for a new report and save it to database
	if weatherReport.ReadReport() != nil {
		weatherReportRaw, err := c.WeatherApp.AskToExternalServiceForWeather(city, country)
		if err != nil {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
			return
		}
		weatherReport.Parse(weatherReportRaw)
		log.Println(weatherReport.WriteReport())
	}

	// Ecode response in json format and write
	c.Ctx.ResponseWriter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Ctx.ResponseWriter)
	encoder.Encode(weatherReport)
}
