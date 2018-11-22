package main

import (
	_"github.com/gislu/gochat/src/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/gochat", "static")
	beego.Run()
}
