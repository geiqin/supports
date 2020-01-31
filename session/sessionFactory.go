package session

import (
	"fmt"
	"github.com/geiqin/supports/config"
	"log"
	"sync"
)

var globalSessionId string
var onceSession sync.Once
var globalSessionManager *Manager
//var globalSession Session

func init()  {
	log.Println("do session init")
	myConf :=&SessionConfig{}
	config.ConvertStruct("session",myConf)
	log.Println(myConf)

	SessionInit(myConf)
}


func SessionInit(cfg * SessionConfig) {
	var err error
	globalSessionManager, err = NewSessionManager(cfg.Provider, cfg.CookieName, 3600)
	if err != nil {
		fmt.Println(err)
		return
	}
	go globalSessionManager.GC()
	fmt.Println("fd")
}

func SessionStart(sessionId string) *error {
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

func SessionDestroy() {
	globalSessionManager.SessionDestroy(globalSessionId)
}

