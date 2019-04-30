package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/poc/logging-service/mediators"
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

	responseData, err := echoMediator(body)
	if err != nil {
		panic(err)
	}

	e.Ctx.Output.SetStatus(http.StatusOK)
	e.Data["json"] = responseData.Message
	e.ServeJSON()
}

func echoMediator(bodyDict map[string]interface{}) (mediators.Response, error){
	mediator := mediators.NewEchoMediator(bodyDict)

	return *mediator, nil
}
