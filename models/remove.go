package models

import (
	"bufio"
	"github.com/astaxie/beego"
	"io"
	"os"
	"os/exec"
	"strings"
)

//吊销证书
func RemoveUser(userid string) {

	cmd := exec.Command("/etc/openvpn/client/easy-rsa/3/remove.sh", userid)
	byte, err := cmd.CombinedOutput()
	if err != nil {
		beego.Error("执行吊销证书脚本错误：", err)
		return
	}
	beego.Info(string(byte))
}

//判断是否存在用户是否吊销
func DelUser(input string) {
	//读取文件中的用户
	file, err := os.Open("/etc/openvpn/easy-rsa/3/pki/index.txt")
	if err != nil {
		beego.Error("open file failed, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				linebyte := strings.Fields(strings.TrimSpace(line))
				name := strings.Split(linebyte[len(linebyte)-1], "/CN=")
				if name[1] == input && linebyte[0] == "V" {
					RemoveUser(input)
					return
				}
			}
			break
		}
		if err != nil {
			beego.Error("read file failed, err:", err)
			return
		}
		linebyte := strings.Fields(strings.TrimSpace(line))
		name := strings.Split(linebyte[len(linebyte)-1], "/CN=")
		if name[1] == input && linebyte[0] == "V" {
			RemoveUser(input)
			return
		}

	}
}
