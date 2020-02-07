package database

import (
	"fmt"
	"github.com/geiqin/supports/config"
	"github.com/geiqin/supports/helper"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

// define our own host type
type DatabaseConfig struct {
	Driver      string    `json:"drviver"`
	Host        string    `json:"host"`
	Port        string    `json:"port"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Database    string    `json:"database"`
	Prefix      string    `json:"prefix"`
	Charset     string    `json:"charset"`
}

var db *gorm.DB


func Load() {

	dbConfig :=&DatabaseConfig{}
	dbcnf :=config.GetConfig("database")
	conns :=config.GetConfig("database","connections")
	defkey,ok :=dbcnf["default"]
	if !ok{
		log.Println("unset defualt value of database config")
		return
	}
	currDb,ok:=conns[defkey.(string)]
	if !ok{
		log.Println("not find defualt database connection")
		return
	}
	helper.MapToStruct(currDb,dbConfig)
	db = createMysqlDB(dbConfig)
	log.Println("load database config succeed")
	if dbConfig.Prefix !="" {
		setDbPrefix(dbConfig.Prefix)
	}
}


func GetDatabase() *gorm.DB {
	if db==nil{
		log.Fatal("not init database , please do InitDatabase function")
	}
	return db
}

/**
设置默认表名前缀
 */
func setDbPrefix(prefix string)  {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return prefix + defaultTableName
	}
}

func createMysqlDB(cfg *DatabaseConfig) *gorm.DB {
	serverAddr :=cfg.Host+":"+cfg.Port
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",cfg.Username, cfg.Password, serverAddr, cfg.Database)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		log.Fatal("Mysql database connection failed")
	}
	return db
}
