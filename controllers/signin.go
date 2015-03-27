package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"luck-web/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplNames = "login.tpl"
}
func (c *LoginController) Post() {
	c.TplNames = "login.tpl"
	c.Ctx.Request.ParseForm()
	email := c.Ctx.Request.Form.Get("email")
	password := c.Ctx.Request.Form.Get("password")
	fmt.Printf("%s\n", email)
	fmt.Printf("%s\n", password)
	isSignuped, key := models.CheckUser(email, password)
	if !isSignuped {
		fmt.Println("user not signuped")
		return
	}
	c.SetSession("email", email)
	c.SetSession("key", key)
	fmt.Printf("email,s %s,key.s %s \n", c.GetSession("email"), c.GetSession("key"))
	c.Redirect("/index", 302)
}
