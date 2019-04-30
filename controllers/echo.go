package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
)

type EchoController struct {
	beego.Controller
}

func (e *EchoController) Echo() {
	body := make(map[string]interface{})

	err := json.Unmarshal(e.Ctx.Input.RequestBody, &body)
	if err != nil {
		panic(err)
	}

	responseData := fmt.Sprintf("Hello %s!", body["name"])

	// TODO: Invocar al mediator

	e.Ctx.Output.SetStatus(http.StatusOK)
	e.Data["json"] = responseData
	e.ServeJSON()
}
