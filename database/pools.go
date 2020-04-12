package database

import (
	"github.com/geiqin/supports/xconfig"
	"github.com/jinzhu/gorm"
	"log"
)

var pools map[string]*gorm.DB
var poolIndex []string

//数据库访问池子
func DbPools(cfg *xconfig.DatabaseInfo, max ...int) *gorm.DB {
	maxLen := 10
	if max != nil {
		maxLen = max[0]
	}
	if pools == nil {
		pools = make(map[string]*gorm.DB)
		poolIndex = make([]string, 0)
	}
	db, ok := pools[cfg.Database]
	if !ok {
		db = CreateMysqlDB(cfg)
		pools[cfg.Database] = db
		poolIndex = append(poolIndex, cfg.Database)
	}
	if len(poolIndex) > maxLen {
		log.Println("db pool is fulled :", len(poolIndex))
	}
	return db
}
