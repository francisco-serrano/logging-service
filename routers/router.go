package routers

import (
	"github.com/astaxie/beego"
	"github.com/poc/logging-service/controllers"
	"github.com/poc/logging-service/filters"
)

func init() {
	ns := beego.NewNamespace("/logging-service",
		beego.NSRouter("/echo", &controllers.EchoController{}, "post:Echo"),
	)

	beego.AddNamespace(ns)

	InitializeFilters()
}

func InitializeFilters() {
	beego.InsertFilter("/logging-service/*", beego.BeforeRouter, filters.InsertMetadata)
	beego.InsertFilter("/logging-service/*", beego.AfterExec, filters.LogRequestMetadata, false)
}