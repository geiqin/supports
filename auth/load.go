package auth

import (
	"github.com/geiqin/supports/xconfig"
	"log"
)

var storeConf *xconfig.TokenInfo
var userConf *xconfig.TokenInfo

func Load() {
	storeConf = xconfig.GetTokenCfg("store")
	if storeConf == nil {
		log.Println("load store_token config failed")
		return
	}
	log.Println("load store_token config succeed")

	userConf = xconfig.GetTokenCfg("user")
	if userConf == nil {
		log.Println("load user_token config failed")
		return
	}
	log.Println("load user_token config succeed")
}
