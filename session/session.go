package session

import (
	"time"
)

var pder Provider

//Session操作接口
type Session interface {
	Set(key string, value interface{}) error
	Get(key string) interface{}
	Delete(key string) error
	Has(key string) bool
	Save() error
	SessionID() string
}

//session实现
type SessionStore struct {
	sid              string                 //session id 唯一标示
	LastAccessedTime time.Time              //最后访问时间
	value            map[string]interface{} //session 里面存储的值
}

//设置
func (st *SessionStore) Set(key string, value interface{}) error {
	st.value[key] = value
	return nil
}

//判断KEY是否存在
func (st *SessionStore) Has(key string) bool {
	_, ok := st.value[key]
	return ok
}

//获取session
func (st *SessionStore) Get(key string) interface{} {
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
	return nil
}

//删除
func (st *SessionStore) Delete(key string) error {
	delete(st.value, key)
	return nil
}

//保存
func (st *SessionStore) Save() error {
	pder.SessionSave(st.sid, st.value)
	return nil
}

func (st *SessionStore) SessionID() string {
	return st.sid
}
