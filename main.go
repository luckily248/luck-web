package main

import (
	"github.com/astaxie/beego"
	_ "luck-web/routers"
)

func main() {
	beego.SessionOn = true
	beego.Run()
}
