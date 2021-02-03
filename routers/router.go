package routers

import (
	"github.com/astaxie/beego"
	"vpnweb/controllers"
)

func init() {
	beego.Router("/", &controllers.Registercontroller{}, "get:GetRegister")
	beego.Router("/register", &controllers.Registercontroller{}, "get:GetRegister2;post:PostRegister")
	beego.Router("/auth", &controllers.Registercontroller{}, "get:GetAuth;post:PostAuth")
}
