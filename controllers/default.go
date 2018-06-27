package controllers

import (
	"beego/orm"
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

	// Create empty response object
	weatherParsed := models.WeatherParsed{
		CodeName: strings.ToLower(city) + strings.ToLower(country),
	}

	o := orm.NewOrm()
	o.Using("default")

	qs := o.QueryTable("reports")
	err := qs.Filter("code_name", weatherParsed.CodeName).One(&weatherParsed)

	if err != nil {
		log.Printf("Record %s NOTexist, getting data from externar service", weatherParsed.CodeName)
		weatherRespose, err := c.WeatherApp.AskToExternalServiceForWeather(city, country)
		if err != nil {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
			return
		}
		weatherParsed.Parse(weatherRespose)
		log.Println(o.Insert(&weatherParsed))
	} else {
		log.Println("Yeah, record exist!")
	}

	c.Ctx.ResponseWriter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Ctx.ResponseWriter)
	err = encoder.Encode(weatherParsed)
}
