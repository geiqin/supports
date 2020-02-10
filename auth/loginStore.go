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

//店铺是否已授权
func StoreAuthed() bool {
	if currentStore==nil{
		return false
	}
	return currentStore.HasLogin
}

//获得当前登录店铺
func GetStore() *LoginStore {
	return currentStore
}

//获得当前登录店铺ID
func GetStoreId() int64 {
	return currentStore.Id
}

//店铺授权
func StoreAuthorization(myStore *LoginStore) *LoginStore {
	onceStore.Do(func() {
		currentStore = myStore
		currentStore.HasLogin =true
	})
	return currentStore
}
