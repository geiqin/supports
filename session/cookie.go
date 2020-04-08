package session

import (
	"sync"
	"time"
)

var httpCookie *HttpCookie
var once sync.Once

/*
HttpCookie结构
MaxAge和Expires都可以设置cookie持久化时的过期时长，Expires是老式的过期方法，
如果可以，应该使用MaxAge设置过期时间，但有些老版本的浏览器不支持MaxAge。
如果要支持所有浏览器，要么使用Expires，要么同时使用MaxAge和Expires。
*/
type  HttpCookie struct{
	Name string
	Value string
	Path string
	HttpOnly bool
	MaxAge   int
	Expires time.Time
}

func GetCookie(cookieName string) *HttpCookie {
	once.Do(func() {
		httpCookie = &HttpCookie{}
	})
	return httpCookie
}

func SetCookie(cookie *HttpCookie) error {
	httpCookie = cookie
	return nil
}

