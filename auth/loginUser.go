package auth

import "github.com/geiqin/supports/session"

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
	val :=sess.Get("user_id")
	if val !=""{
		return val.(int64)
	}
	return 0
}

//用户授权
func UserAuthorization(authUser *LoginUser) *LoginUser {
	currentUser = authUser
	currentUser.HasLogin =true
	return currentUser
}
