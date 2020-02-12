package main

import (
	"github.com/geiqin/supports/app"
	"github.com/geiqin/supports/auth"
	"github.com/geiqin/supports/cache"
	"github.com/geiqin/supports/helper"
	"github.com/geiqin/supports/session"
	"github.com/geiqin/supports/token"
	log "log"
)


//生成用户Token
func MakeUserToken(user *auth.LoginUser,clientIp string) (string,error) {
	accKey :=helper.UniqueId()
	token :=token.UserToken{}
	t, err :=token.Encode(user,&auth.AccessLimit{
		AccessKey: accKey,
		ClientIp:  clientIp,
	})
	if err != nil {
		return "",err
	}
	return t,nil
}

func main() {

	app.Run("srv_cms_media")



	myCh :=cache.GetCache()
	myCh.Set("ddd","123",0)
	myCh.Get("ddd")
	/*

	//clientIp :=lib.GetIP(ctx)
	clientIp :="127.0.0.1:8688"
	log.Println("client ip:",clientIp)

	token, err :=MakeUserToken(&auth.LoginUser{
		Id:  3,
		Name: "aaa",
	},clientIp)

	//log.Println("token:",token)
	//log.Println("err:",err)
	if err ==nil{
		log.Println("token:",token)
	}

	 */


	session.Start("ppVYPYbn1rq2H_yK7fUZyuhTC1LKshCH2cr0Jbn_fwo=")

	ss:=session.GetSession()

	//log.Println("session:",ss)

	log.Println("session id:",ss.SessionID())
	ss.Set("name","zhangshan")
	ss.Set("sex","man")
	log.Println("session name:",ss.Get("name"))
	log.Println("session sex:",ss.Get("sex"))
	//ss.
	//log.Println("session id:",ss.SessionID())
}