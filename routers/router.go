package routers

import (
	"github.com/astaxie/beego"
	"github.com/ewangsong/vpnweb/controllers"
)

func init() {
	beego.Router("/", &controllers.Registercontroller{}, "get:GetRegister;post:PostRegister")
}
