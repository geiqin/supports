package database

import (
	"fmt"
	"github.com/geiqin/supports/config"
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

func init() {
	log.Println("do database init")
	dbConfig :=&DatabaseConfig{}
	config.ConvertStruct("mysql",dbConfig)
	db = createMysqlDB(dbConfig)
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
