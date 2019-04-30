package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/poc/logging-service/routers"
	//log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("iniciando logging-services")

	//log.WithFields(log.Fields{
	//	"animal": "walrus",
	//}).Info("A walrus appears")

	beego.Run()
}
