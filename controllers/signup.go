package controllers

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"luck-web/models"
	"luck-web/utils"
	"strings"
)

type SignUpController struct {
	beego.Controller
}

func (c *SignUpController) Get() {
	c.TplNames = "signup.tpl"
}
func (c *SignUpController) Post() {
	c.TplNames = "signup.tpl"
	c.Ctx.Request.ParseForm()
	email := c.Ctx.Request.Form.Get("email")
	password := c.Ctx.Request.Form.Get("password")
	salt := strings.Replace(uuid.NewUUID().String(), "-", "", -1)
	newkey := utils.GetKey(password, salt)
	user := models.User{}
	user.Email = email
	user.Password = newkey
	user.Salt = salt
	fmt.Printf("%s\n", email)
	fmt.Printf("%s\n", newkey)
	err := models.AddUser(user)
	if err != nil {
		log.Fatal("add error")
	}
	c.Redirect("/", 302)
}
