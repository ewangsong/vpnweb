package controllers

import "github.com/astaxie/beego"

type Registercontroller struct {
	beego.Controller
}

func (r *Registercontroller) GetRegister() {
	r.TplName = "index.tpl"
}
func (r *Registercontroller) PostRegister() {

}
