package auth

import (
	"sync"
)

var currentStore *CurrentStore
var onceStore sync.Once

type CurrentStore struct {
	Id   int32
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

func GetCurrentStore() *CurrentStore {
	return currentStore
}

func InitCurrentStore(initStore *CurrentStore) *CurrentStore {
	onceStore.Do(func() {
		currentStore = initStore
		currentStore.HasLogin =true
	})
	return currentStore
}
