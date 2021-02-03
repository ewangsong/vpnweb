package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "vpnweb/routers"
)

func init() {
	beego.AppPath = "/opt/vpnweb"
	beego.SetViewsPath("/opt/vpnweb/views/")
	beego.LoadAppConfig("ini", "/opt/vpnweb/conf/app.conf")
	jsonConfig := `{
	    "filename" : "/var/log/vpnweb/vpnweb.log"
	}` //定义日志文件路径和名字
	// jsonConfig := `{
	//     "filename" : "./lanradius.log"
	// }` //定义日志文件路径和名字

	logs.SetLogger(logs.AdapterFile, jsonConfig) // 设置日志记录方式：本地文件记录
	logs.EnableFuncCallDepth(true)               // 输出log时能显示输出文件名和行号（非必须）
	beego.BeeLogger.DelLogger("console")         //删除console日志输出
}
func main() {
	beego.Run()
}
