package auth

import (
	"sync"
)

var currentUser *LoginUser
var onceUser sync.Once

//用户类型
type UserType string

const (
	User UserType = "user"
	Customer UserType ="customer"
)

type AccessLimit struct {
	AccessKey string
	ClientIp string
}

type LoginUser struct {
	Id   int64
	Type  UserType
	HasLogin bool
}

//判断用户是否登录
func UserAuthed() bool {
	if currentUser==nil{
		return false
	}
	return currentUser.HasLogin
}

//获得当前登录用户
func GetUser() *LoginUser {
	onceUser.Do(func() {
		currentUser = &LoginUser{}
	})
	return currentUser
}

//获得当前登录用户ID
func GetUserId() int64 {
	return currentStore.Id
}

//用户授权
func UserAuthorization(authUser *LoginUser) *LoginUser {
	onceUser.Do(func() {
		currentUser = authUser
		currentUser.HasLogin =true
	})
	return currentUser
}
