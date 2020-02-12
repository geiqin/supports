package session

import (
	"time"
)

var pder Provider

//session实现
type SessionStore struct {
	sid              string                      //session id 唯一标示
	LastAccessedTime time.Time                   //最后访问时间
	value            map[string]interface{} //session 里面存储的值
}

//设置
func (st *SessionStore) Set(key string, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}

//获取session
func (st *SessionStore) Get(key string) interface{} {
	pder.SessionUpdate(st.sid)
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
	pder.SessionUpdate(st.sid)
	return nil
}

//保存
func (st *SessionStore) Save() error {
	pder.SessionSave(st.sid,st.value)
	return nil
}


func (st *SessionStore) SessionID() string {
	return st.sid
}