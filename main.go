package main

import (
	"github.com/astaxie/beego"
	"github.com/poc/logging-service/controllers"
	_ "github.com/poc/logging-service/routers"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
