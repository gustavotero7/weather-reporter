package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (er *ErrorController) Error404() {

	er.Ctx.ResponseWriter.Write([]byte("Meh, i'm a 404 error, go away dude. v2.0"))
}
func (er *ErrorController) Error500() {

}
