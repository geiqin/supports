package main

import (
	"context"
	"github.com/geiqin/supports/app"
	"github.com/geiqin/supports/auth"
	"github.com/geiqin/supports/cache"
	"github.com/geiqin/supports/database"
	"github.com/geiqin/supports/helper"
	"github.com/geiqin/supports/session"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/metadata"
	log "log"
)

//生成用户Token
func MakeUserToken(user *auth.LoginUser, clientIp string) (string, error) {
	accKey := helper.UniqueId()
	token := auth.UserToken{}
	t, err := token.Encode(user, &auth.AccessLimit{
		AccessKey: accKey,
		ClientIp:  clientIp,
	})
	if err != nil {
		return "", err
	}
	return t, nil
}

func ConnDB() *gorm.DB {
	cfg := database.GetDbCfg()
	db := database.DbPools(cfg, 10)
	return db
}

func main222() {
	log.Println("code:", helper.GenerateSn())
	log.Println("code:", helper.GenerateSn("2018"))
	app.Run("srv_dms", true)
	db := ConnDB()
	log.Println(db)
}

func NewContext(headers map[string]string) context.Context {
	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), headers)
	return ctx
}

func main() {

	app.Run("srv_supports", true)

	//log.Println("pinyin:", helper.ConvertPinyin("我是中国人"))
	myCh := cache.GetCache()
	log.Println("cache value:", myCh.Get("ddd"))
	myCh.Get("ddd")
	myCh.Set("ddd", "1211113", 0)
	log.Println("cache value2:", myCh.Get("ddd"))

	//log.Println("cache key:",myCh.Get("storekey"))

	ctx := NewContext(map[string]string{
		"Session-Id": "xZNo_6ulP6xE9aXQ6TWO0n75lAgpi34aqQnUPEDKeTQ",
	})

	ss := session.GetSession(ctx)
	log.Println("hash value:", ss.Get("hash"))
	ss.Set("hash", "555555555555")
	//ss.Set("key","aaaaaaa")
	//ss.Save()
	//va,_ :=xtime.ParseTimeByTimeStr("2019-12-11 12:10:30","d")

	//log.Println("time  :",va)
	ss.Save()
	//log.Println("session:",ss)
	log.Println("user key:", ss.Get("hash"))

	log.Println("session id:", ss.SessionID())

	log.Println("session id:", ss.SessionID())
	log.Println("user id:", ss.Get("user_id"))
	log.Println("user name:", ss.Get("user_name"))
	log.Println("user_mobile:", ss.Get("user_mobile"))
	//ss.
	//log.Println("session id:",ss.SessionID())
}
