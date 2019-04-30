package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/pkg/errors"
	"github.com/poc/logging-service/utils"
)

type ErrorController struct {
	beego.Controller
	RC utils.ResponseComposer
}

func (e *ErrorController) Error404() {
	err := utils.ErrorNotFound(errors.New("invalid URL"))
	e.RC.SetErrorResponse(&e.Controller, utils.GetStatusErrorCode(err), err)
	fmt.Println(err)

	e.ServeJSON()
}
