package token

import (
	"github.com/geiqin/supports/config"
	"log"
)

var storeConf *ConfToken
var userConf *ConfToken

type ConfToken struct {
	Issuer string `json:"issuer"`
	Audience string `json:"audience"`
	PrivateKey []byte `json:"private_key"`
	ExpireTime int64 `json:"expire_time"`
}

func Load() {
	conf :=config.GetConfig("auth","providers")
	if conf ==nil{
		log.Println("load token config failed")
	}
	log.Println("load token config succeed")

	storeConfObj,ok :=conf["store"]
	if !ok{
		log.Println("load store token config failed")
	}
	ToConfToken(storeConf,storeConfObj.(map[string]interface{}))
	log.Println("load store token config succeed")

	userConfObj,ok :=conf["user"]
	if !ok{
		log.Println("load user token config failed")
	}
	ToConfToken(userConf,userConfObj.(map[string]interface{}))
	log.Println("load user token config succeed")
}


func ToConfToken(to *ConfToken,from map[string]interface{}) *ConfToken{

	to =&ConfToken{
		Issuer: from["issuer"].(string),
		Audience: from["audience"].(string),
		PrivateKey: []byte(from["private_key"].(string)),
		//ExpireTime: from["expire_time"].(int64),
	}
	return to
}