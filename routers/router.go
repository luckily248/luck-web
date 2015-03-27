package routers

import (
	"github.com/astaxie/beego"
	"luck-web/controllers"
)

func init() {
	//beego.Router("/logincheck", &controllers.LoginCheckController{})
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/signup", &controllers.SignUpController{})
	beego.Router("/index", &controllers.MainController{})
}
