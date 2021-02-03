//发送邮件验证
package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/gomail.v2"
	"strconv"
)

//发送邮件
func SendMail(mailTo string, name string, id string) error {
	url := beego.AppConfig.String("registerurl")
	subject := "openVPN账号注册"

	url = url + "/auth?email=" + mailTo + "&RegisterID=" + id
	body := "Hello, " + name + " 这是一封注册邮件，请点击<a href=" + url + ">此处</a>完成注册," + "如果点击无效请将地址拷贝到浏览器地址栏直接" +
		"访问完成确认" + "<br>" + "<a href=" + url + ">" + url + "</a>"
	//邮件信息
	mailConn := map[string]string{
		"user": "wangsong1@wsecar.com",
		"pass": "Ufenqi@246",
		"host": "smtphz.qiye.163.com",
		"port": "25",
	}
	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "openVPN")) //这种方式可以添加别名，即“XX官方”
	//说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])

	m.SetHeader("To", mailTo)       //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err
}
