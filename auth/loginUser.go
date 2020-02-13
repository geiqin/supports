package auth

import (
	"github.com/geiqin/supports/helper"
	"github.com/geiqin/supports/session"
	"log"
)

var currentUser *LoginUser
//var onceUser sync.Once


type AccessLimit struct {
	AccessKey string
	ClientIp string
}

type LoginUser struct {
	Id   int64
	Type  string
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
	if currentUser ==nil{
		currentUser = &LoginUser{}
	}
	return currentUser
}

//获得当前登录用户ID
func GetUserId() int64 {
	sess :=session.GetSession()
	vals :=sess.Get("user_id")
	log.Println("auth_getuserid:",vals)
	val :=helper.StringToInt64(vals.(string))
	 
	return val
}

//用户授权
func UserAuthorization(authUser *LoginUser) *LoginUser {
	currentUser = authUser
	currentUser.HasLogin =true
	return currentUser
}
