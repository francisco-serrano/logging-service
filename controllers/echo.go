package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
)

type EchoController struct {
	beego.Controller
}

func (e *EchoController) Echo() {
	responseData := "hola como andas"

	// TODO: Invocar al mediator

	println("EchoController")

	e.Ctx.Output.SetStatus(http.StatusOK)
	e.Data["json"] = responseData
	e.ServeJSON()
}
