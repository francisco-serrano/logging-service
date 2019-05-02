package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/poc/logging-service/filters"
	"github.com/poc/logging-service/utils"
)

type ErrorController struct {
	beego.Controller
	RC utils.ResponseComposer
}

type Mensaje struct {
	Id        int
	Contenido string
}

func (e *ErrorController) Error404() {
	defer e.ServeJSON()

	err := utils.ErrorNotFound(errors.New("invalid URL"))

	e.EnableRender = false
	e.RC.SetErrorResponse(&e.Controller, utils.GetStatusErrorCode(err), err)

	filters.LogRequestMetadata(e.Ctx)
}
