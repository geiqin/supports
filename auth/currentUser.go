package auth

import (
	"sync"
)

var currentUser *CurrentUser
var onceUser sync.Once


type AccessLimit struct {
	AccessKey string
	ClientIp string
}

type CurrentUser struct {
	Id   int32
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

func GetCurrentUser() *CurrentUser {
	onceUser.Do(func() {
		currentUser = &CurrentUser{}
	})
	return currentUser
}

func InitCurrentUser(initUser *CurrentUser) *CurrentUser {
	onceUser.Do(func() {
		currentUser = initUser
		currentUser.HasLogin =true
	})
	return currentUser
}
