package database

import (
	"github.com/geiqin/supports/xconfig"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

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
	//parseTime=True&    /utf8mb4    Local&time_zone=   Asia%2FShanghai

	//当前有效两种
	connString := cfg.Username + ":" + cfg.Password + "@tcp(" + serverAddr + ")/" + cfg.Database + "?charset=utf8mb4&loc=Local"
	//connString := cfg.Username + ":" + cfg.Password + "@tcp(" + serverAddr + ")/" + cfg.Database + "?charset=utf8mb4&loc=Local&parseTime=True"

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
