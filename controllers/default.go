package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Post() {
	c.Data["text"] = "beego.me"
	c.TplNames = "index.tpl"
}
func (c *MainController) Get() {
	c.Data["text"] = "beego.me"
	c.TplNames = "index.tpl"
}
