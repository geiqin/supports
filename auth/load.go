package auth

import (
	"github.com/geiqin/supports/config"
	"log"
	"strconv"
)

var storeConf *ConfToken
var userConf *ConfToken
var customerConf *ConfToken

type ConfToken struct {
	Issuer string `json:"issuer"`
	Audience string `json:"audience"`
	PrivateKey []byte `json:"private_key"`
	ExpireTime int `json:"expire_time"`
}

func Load() {
	conf :=config.GetConfig("auth","providers")
	if conf ==nil{
		log.Println("load tokens config failed")
		return
	}
	log.Println("load tokens config succeed")

	storeConfObj,ok :=conf["store"]
	if !ok{
		log.Println("load store_token config failed")
		return
	}
	storeConf =ToConfToken(storeConf,storeConfObj.(map[string]interface{}))
	log.Println("load store_token config succeed")

	userConfObj,ok :=conf["user"]
	if !ok{
		log.Println("load user_token config failed")
		return
	}
	userConf =ToConfToken(userConf,userConfObj.(map[string]interface{}))
	log.Println("load user_token config succeed")

}


func ToConfToken(to *ConfToken,from map[string]interface{}) *ConfToken{
	exTime,err :=strconv.Atoi(from["expire_time"].(string))
	if err!=nil{
		exTime =15
	}
	to =&ConfToken{
		Issuer: from["issuer"].(string),
		Audience: from["audience"].(string),
		PrivateKey: []byte(from["private_key"].(string)),
		ExpireTime: exTime,
	}
	return to
}