package models

import (
	"math/rand"
	"time"
)

//创建随机字符串，作为申请注册的验证码
func CreateRandomString() string {
	s := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	id := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 25; i++ {
		n := rand.Intn(62)
		id += string(s[n])
	}
	return id
}
