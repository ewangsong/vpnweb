package routers

import (
	"github.com/astaxie/beego"
	"vpnweb/controllers"
)

func init() {
	beego.Router("/", &controllers.Registercontroller{}, "get:GetRegister")
	beego.Router("/register", &controllers.Registercontroller{}, "get:GetRegister;post:PostRegister")
}
