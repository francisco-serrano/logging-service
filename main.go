package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/poc/logging-service/routers"
)

func main() {
	fmt.Println("hola logging-service")

	beego.Run()
}
