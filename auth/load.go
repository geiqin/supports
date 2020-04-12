package auth

import (
	"github.com/geiqin/supports/xconfig"
	"log"
)

var storeConf *xconfig.TokenInfo
var userConf *xconfig.TokenInfo

func Load() {
	storeCfg := xconfig.GetTokenCfg("store")
	if storeCfg == nil {
		log.Println("load store_token config failed")
		return
	}
	log.Println("load store_token config succeed")

	userCfg := xconfig.GetTokenCfg("user")
	if userCfg == nil {
		log.Println("load user_token config failed")
		return
	}
	log.Println("load user_token config succeed")
}
