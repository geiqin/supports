package auth

import (
	"github.com/geiqin/supports/xconfig"
	"log"
)

var storeConf *xconfig.AuthInfo
var userConf *xconfig.AuthInfo

func Load() {
	storeCfg := xconfig.GetAuthCfg("store")
	if storeCfg == nil {
		log.Println("load store_token config failed")
		return
	}
	log.Println("load store_token config succeed")

	userCfg := xconfig.GetAuthCfg("user")
	if userCfg == nil {
		log.Println("load user_token config failed")
		return
	}
	log.Println("load user_token config succeed")
}
