package auth

import (
	"sync"
)

var currentUser *LoginUser
var onceUser sync.Once


type AccessLimit struct {
	AccessKey string
	ClientIp string
}

type LoginUser struct {
	Id   int64
	Name  string
	HasLogin bool
}

//判断用户是否登录
func HasUserLogin() bool {
	if currentUser==nil{
		return false
	}
	return currentUser.HasLogin
}

//获得当前登录用户
func GetLoginUser() *LoginUser {
	onceUser.Do(func() {
		currentUser = &LoginUser{}
	})
	return currentUser
}

func InitLoginUser(initUser *LoginUser) *LoginUser {
	onceUser.Do(func() {
		currentUser = initUser
		currentUser.HasLogin =true
	})
	return currentUser
}
