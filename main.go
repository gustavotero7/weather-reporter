package main

import (
	"beego/orm"
	"log"
	_ "weather-reporter/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	log.Println("Initializing db")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:password@/weatherreporter?charset=utf8")
}

func main() {

	//beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
