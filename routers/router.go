package routers

import (
	"beego_project/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "post:Post")
	beego.Router("/signup", &controllers.SignUpController{})
}
