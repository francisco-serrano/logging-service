package routers

import (
	"github.com/astaxie/beego"
	"github.com/poc/logging-service/controllers"
)

func init() {
	//beego.Router("/echo", &controllers.EchoController{}, "get:Echo")

	ns := beego.NewNamespace("/logging-service",
		beego.NSRouter("/echo", &controllers.EchoController{}, "post:Echo"),
	)

	beego.AddNamespace(ns)
}