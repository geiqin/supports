package main

import (
	"github.com/geiqin/supports/app"
	"github.com/geiqin/supports/auth"
	"github.com/geiqin/supports/database"
	"github.com/geiqin/supports/helper"
	"github.com/geiqin/supports/session"
	"github.com/jinzhu/gorm"
	log "log"
)


//生成用户Token
func MakeUserToken(user *auth.LoginUser,clientIp string) (string,error) {
	accKey :=helper.UniqueId()
	token := auth.UserToken{}
	t, err :=token.Encode(user,&auth.AccessLimit{
		AccessKey: accKey,
		ClientIp:  clientIp,
	})
	if err != nil {
		return "",err
	}
	return t,nil
}


func ConnDB() *gorm.DB {
	cfg := database.GetDbCfg()
	return database.CreateMysqlDB(cfg)
}


func main() {
	log.Println("code:", helper.GenerateSn())
	log.Println("code:", helper.GenerateSn("2018"))
	app.Run("srv_dms")
	db :=ConnDB()
	log.Println(db)
}

func main222() {

	app.Run("srv_supports")

	//myCh :=cache.GetCache()
	//myCh.Set("ddd","1211113",0)
	//myCh.Get("ddd")

	//log.Println("cache key:",myCh.Get("storekey"))

	session.Start("xZNo_6ulP6xE9aXQ6TWO0n75lAgpi34aqQnUPEDKeTQ=")

	ss:=session.GetSession()
	ss.Set("hash","555555555555")
	//ss.Set("key","aaaaaaa")
	//ss.Save()
	//va,_ :=xtime.ParseTimeByTimeStr("2019-12-11 12:10:30","d")

	//log.Println("time  :",va)

	//log.Println("session:",ss)
	log.Println("user key:",ss.Get("hash"))


	log.Println("session id:",ss.SessionID())


	log.Println("session id:",ss.SessionID())
	log.Println("user id:",ss.Get("user_id"))
	log.Println("user name:",ss.Get("user_name"))
	log.Println("user_mobile:",ss.Get("user_mobile"))
	//ss.
	//log.Println("session id:",ss.SessionID())
}