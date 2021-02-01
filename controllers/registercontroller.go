package controllers

import (
	"github.com/astaxie/beego"
)

type Registercontroller struct {
	beego.Controller
}

func (r *Registercontroller) GetRegister() {
	r.TplName = "index.html"
}
func (r *Registercontroller) PostRegister() {

}
