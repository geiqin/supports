package session

import (
	"context"
	"fmt"
	"github.com/geiqin/supports/config"
	"github.com/geiqin/supports/helper"
	"github.com/micro/go-micro/v2/metadata"
	"log"
)

var globalSessionManager *Manager

func Load() {
	myConf := &SessionConfig{}
	cnf := config.GetConfig("app", "session")
	if cnf == nil {
		log.Println("load session config failed")
	}
	helper.MapToStruct(cnf, myConf)
	log.Println("load session config succeed")
	LoadRedis(myConf)
	newManager(myConf)
}

func newManager(cfg *SessionConfig) {
	var err error
	globalSessionManager, err = NewSessionManager(cfg.Provider, cfg.CookieName, 3600)
	if err != nil {
		fmt.Println(err)
		return
	}
	if cfg.Provider == "memory" {
		go globalSessionManager.GC()
	}

	//fmt.Println("session ok")
}

func GetSession(ctx context.Context) (session Session) {
	session = globalSessionManager.SessionStart(GetSessionId(ctx))
	return session
}

func GetSessionById(sessionId string) (session Session) {
	session = globalSessionManager.SessionStart(sessionId)
	return session
}

func Destroy(ctx context.Context) {
	globalSessionManager.SessionDestroy(GetSessionId(ctx))
}

func GetSessionId(ctx context.Context) string {
	meta, ok := metadata.FromContext(ctx)
	if ok {
		return meta["Session-Id"]
	}
	return ""
}
