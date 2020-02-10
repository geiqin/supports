package auth

import (
	"sync"
)

var currentCustomer *LoginCustomer
var onceCustomer sync.Once

type LoginCustomer struct {
	Id   int64
	Name  string
	HasLogin bool
}

//判断客户是否登录
func CustomerAuthed() bool {
	if currentUser==nil{
		return false
	}
	return currentUser.HasLogin
}

//获得当前登录客户
func GetCustomer() *LoginCustomer {
	onceCustomer.Do(func() {
		currentCustomer = &LoginCustomer{}
	})
	return currentCustomer
}

//获得当前登录客户ID
func GetCustomerId() int64 {
	return currentCustomer.Id
}

//客户授权
func CustomerAuthorization(authCustomer *LoginCustomer) *LoginCustomer {
	onceCustomer.Do(func() {
		currentCustomer = authCustomer
		currentCustomer.HasLogin =true
	})
	return currentCustomer
}
