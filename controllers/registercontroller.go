package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
	"vpnweb/models"
)

type Registercontroller struct {
	beego.Controller
}

//访问首页
func (r *Registercontroller) GetRegister() {
	r.Redirect("/register.html", 302)
}

//注册页面
func (r *Registercontroller) GetRegister2() {
	r.TplName = "register.html"
}

//注册页面post方法
func (r *Registercontroller) PostRegister() {
	email := r.GetString("email")
	name := r.GetString("realname")
	ok, err := models.CheckUser(cutemail(email))
	if err != nil {
		beego.Error("打开index.txt文件错误", err)
		r.Ctx.WriteString(fmt.Sprint(err))
		return
	}
	if ok {
		r.Ctx.WriteString(`<html>
		<body>
		<h2 style="background-color:red">此用户已注册，注册邮件取消发送，如有误请联系管理员！！</h2>
		</body>
		</html>`)
		return
	}
	id := models.CreateRandomString()
	err = models.SendMail(email, name, id)
	if err != nil {
		beego.Error("注册页面发送邮件失败，失败信息如下：", err)
		return
	}

	//设置session
	r.SetSession(id, email)
	r.TplName = "register.html"
}

//验证页面
func (r *Registercontroller) GetAuth() {
	r.TplName = "auth.html"
	registerid := r.Ctx.Input.Query("RegisterID")
	getemail := r.Ctx.Input.Query("email")
	email := r.GetSession(registerid)
	if getemail != email {
		r.Ctx.WriteString("注册失败，请重新注册或联系管理员")
	}
	userid := cutemail(getemail)

	//执行脚本添加用户
	models.AddUser(userid)
}

//验证post
func (r *Registercontroller) PostAuth() {
	r.TplName = "auth.html"
	id := r.Ctx.Input.Query("RegisterID")
	x := r.GetSession(id)
	v, ok := x.(string)
	if ok {
		v = cutemail(v)
		models.ZipFile(v)
		//下载配置好的文件
		r.Ctx.Output.Download("/opt/vpnweb/client/" + v + ".zip")
	}
}

//cutemail 切割邮箱地址获取账号
func cutemail(email string) string {
	email = strings.TrimSuffix(strings.TrimSpace(email), "@wsecar.com")
	return email
}
