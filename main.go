package main

import (
	"github.com/astaxie/beego"
	_ "github.com/ewangsong/vpnweb/models"
	_ "github.com/ewangsong/vpnweb/routers"
)

func main() {
	beego.Run()
}
