package main

import (
	_ "weather-reporter/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

