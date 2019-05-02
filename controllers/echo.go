package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/poc/logging-service/mediators"
	"github.com/poc/logging-service/utils"
	"net/http"
)

type EchoController struct {
	beego.Controller
	RC utils.ResponseComposer
}

func (e *EchoController) Echo() {
	defer e.ServeJSON()

	body := make(map[string]interface{})

	err := json.Unmarshal(e.Ctx.Input.RequestBody, &body)
	if err != nil {
		err := utils.ErrorInvalidData(err)
		e.RC.SetErrorResponse(&e.Controller, utils.GetStatusErrorCode(err), err)
		e.Ctx.Output.SetStatus(http.StatusBadRequest)
		return
	}

	responseData, err := echoMediator(body)
	if err != nil {
		panic(err)
	}

	e.Ctx.Output.SetStatus(http.StatusOK)
	e.Data["json"] = responseData.Message
}

func echoMediator(bodyDict map[string]interface{}) (mediators.Response, error) {
	mediator := mediators.NewEchoMediator(bodyDict)

	return *mediator, nil
}
