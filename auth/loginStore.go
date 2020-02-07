package auth

import (
	"sync"
)

var currentStore *LoginStore
var onceStore sync.Once

type LoginStore struct {
	Id   int64
	Name  string
	HasLogin bool
}

//判断店铺是否登录
func HasStoreLogin() bool {
	if currentStore==nil{
		return false
	}
	return currentStore.HasLogin
}

//获得当前登录店铺
func GetLoginStore() *LoginStore {
	return currentStore
}

func InitLoginStore(initStore *LoginStore) *LoginStore {
	onceStore.Do(func() {
		currentStore = initStore
		currentStore.HasLogin =true
	})
	return currentStore
}
