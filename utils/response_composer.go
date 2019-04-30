package utils

import "github.com/astaxie/beego"

type ResponseComposer struct {

}

type ResponseEnvelope struct {
	Code int
	Data interface{}
}

func (rc *ResponseComposer) ComposeResponse(statusCode int, data interface{}) ResponseEnvelope {
	responseEnvelope := ResponseEnvelope{}
	responseEnvelope.Code = statusCode
	responseEnvelope.Data = data

	return responseEnvelope
}

func (rc *ResponseComposer) SetErrorResponse(c *beego.Controller, statusCode int, data *Error) {
	responseEnvelope := rc.ComposeResponse(statusCode, data)

	c.Ctx.Output.SetStatus(statusCode)
	c.Data["json"] = responseEnvelope
}
