package database

import (
	"github.com/geiqin/supports/xconfig"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

/*
// define our own host type
type DbConfig struct {
	Driver   string `json:"drviver"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Prefix   string `json:"prefix"`
	Charset  string `json:"charset"`
}
*/

//var db *gorm.DB

//var dbConfigs map[string] DbConfig
//var dbConfig *xconfig.DatabaseInfo
/*
func Load(flag string) {
	//dbConfig = &DbConfig{}
	//connCfg := config.GetConfig("database", "connections", flag)
	dbConfig = xconfig.GetDatabaseCfg(flag)
	if dbConfig == nil {
		log.Println("load database config failed")
		return
	}

	//helper.MapToStruct(connCfg, dbConfig)
	db := CreateMysqlDB(dbConfig)
	log.Println("load database config succeed")
	if dbConfig.Prefix != "" {
		setDbPrefix(db, dbConfig.Prefix)
	}
}

func GetDatabase() *gorm.DB {
	if db == nil {
		log.Println("not init database , please do InitDatabase function")
	}
	return db
}


*/
/**
设置默认表名前缀
*/
func setDbPrefix(db *gorm.DB, prefix string) *gorm.DB {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return prefix + defaultTableName
	}
	return db
}

func CreateMysqlDB(cfg *xconfig.DatabaseInfo) *gorm.DB {
	serverAddr := cfg.Host + ":" + cfg.Port
	//timezone := "'Asia/Shanghai'"
	//connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", cfg.Username, cfg.Password, serverAddr, cfg.Database)
	//parseTime=True&    /utf8mb4    Local&time_zone=    Asia%2FShanghai
	connString := cfg.Username + ":" + cfg.Password + "@tcp(" + serverAddr + ")/" + cfg.Database + "?charset=utf8&loc=Local"
	//connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local&time_zone=%27Asia%2FShanghai%27", cfg.Username, cfg.Password, serverAddr, cfg.Database)
	//connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&loc=Local", cfg.Username, cfg.Password, serverAddr, cfg.Database)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		log.Println("mysql database connection failed")
	}

	if cfg.Prefix != "" {
		setDbPrefix(db, cfg.Prefix)
	}

	return db
}

/*
func GetDbCfg(dbName ...string) *xconfig.DatabaseInfo {
	if dbName != nil {
		dbConfig.Database = dbName[0]
	}
	return dbConfig
}
*
*/
