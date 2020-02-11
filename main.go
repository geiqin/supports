package main

import (
	"github.com/geiqin/supports/auth"
	"github.com/geiqin/supports/helper"
	"github.com/geiqin/supports/token"
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

	//app.Run("srv_cms_media")

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



}