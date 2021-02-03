package models

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"os/exec"
	"strings"
)

//添加用户
func AddUser(userid string) {

	cmd := exec.Command("/etc/openvpn/client/easy-rsa/3/adduser.sh", userid)
	byte, err := cmd.CombinedOutput()
	if err != nil {
		beego.Error("执行添加用户证书脚本错误：", err)
		return
	}
	beego.Info(string(byte))
}

//判断是否存在用户
func CheckUser(input string) (ok bool, err error) {
	//读取文件中的用户
	file, err := os.Open("/etc/openvpn/easy-rsa/3/pki/index.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return false, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				linebyte := strings.Split(strings.TrimSpace(line), "/CN=")
				if input == linebyte[1] {
					return true, nil
				}
			}
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return false, err
		}
		linebyte := strings.Split(strings.TrimSpace(line), "/CN=")
		if input == linebyte[1] {
			return true, nil
		}
	}
	return false, nil
}
