package models

import (
	"archive/zip"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ZipFile(name string) {
	src := "/etc/openvpn/client/" + name
	zipFileName := name + ".zip"
	fmt.Println(src, zipFileName)
	Zip(src, zipFileName)

}

// 打包成zip文件
func Zip(srcDir string, zipFileName string) {

	// 预防：旧文件无法覆盖
	os.RemoveAll(zipFileName)

	// 创建：zip文件
	zipfile, _ := os.Create(zipFileName)
	defer zipfile.Close()

	// 打开：zip文件
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// 遍历路径信息
	filepath.Walk(srcDir, func(path string, info os.FileInfo, _ error) error {
		// 如果是源路径，提前进行下一个遍历
		if path == srcDir {
			return nil
		}

		header, _ := zip.FileInfoHeader(info)
		path = strings.ReplaceAll(path, "\\", "/")     // 对同时兼容Linux 和 win 进行处理
		header.Name = strings.TrimPrefix(path, srcDir) // +`/`
		//	beego.Info(header.Name,path,src_dir) // pictures\10.jpg E:xx\V1\pictures\10.jpg E:\xx\V1

		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}

		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})
	err := os.Rename(zipFileName, "/opt/vpnweb/client/"+zipFileName)
	if err != nil {
		beego.Error("压缩完成以后移动错误", err)
	}
}
