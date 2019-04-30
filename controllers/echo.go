package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
)

type EchoController struct {
	beego.Controller
}

// @router /:msg [post]
func (e *EchoController) Echo() {
	msg := e.Ctx.Input.Param(":msg")
	responseData := fmt.Sprintf("Hello %s!", msg)

	// TODO: Invocar al mediator

	println("EchoController -> " + msg)

	e.Ctx.Output.SetStatus(http.StatusOK)
	e.Data["json"] = responseData
	e.ServeJSON()
}
