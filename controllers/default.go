package controllers

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dchest/scrypt"
	"io"
	"time"
)

type MainController struct {
	beego.Controller
}

type LoginController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}

func (c *LoginController) Get() {
	c.TplNames = "login.tpl"
}
func (c *LoginController) Post() {
	c.TplNames = "login.tpl"
	c.Ctx.Request.ParseForm()
	email := c.Ctx.Request.Form.Get("email")
	password := c.Ctx.Request.Form.Get("password")
	salt := "goodluck"
	dk := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)

}
