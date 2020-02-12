package session

import (
	"fmt"
	"github.com/geiqin/supports/config"
	"github.com/geiqin/supports/helper"
	"log"
	"sync"
)

var globalSessionId string
var onceSession sync.Once
var globalSessionManager *Manager
//var globalSession Session

func Load()  {
	myConf :=&SessionConfig{}
	cnf :=config.GetConfig("app","session")
	if cnf ==nil{
		log.Println("load session config failed")
	}
	helper.MapToStruct(cnf,myConf)
	log.Println("load session config succeed")
	LoadRedis(myConf)
	newManager(myConf)
}


func newManager(cfg * SessionConfig) {
	var err error
	globalSessionManager, err = NewSessionManager(cfg.Provider, cfg.CookieName, 3600)
	if err != nil {
		fmt.Println(err)
		return
	}
	if cfg.Provider=="memory"{
		go globalSessionManager.GC()
	}

	//fmt.Println("session ok")
}

func Start(sessionId string) *error {
	onceSession.Do(func() {
		globalSessionId =sessionId
	})
	return nil
}

func GetSession() (session Session){
	//onceSession.Do(func() {
	session =globalSessionManager.SessionStart(globalSessionId)
	//})
	return session
}

func Destroy() {
	globalSessionManager.SessionDestroy(globalSessionId)
}

