package database

import "github.com/jinzhu/gorm"

var pools map[string]*gorm.DB
var poolIndex []string

//数据库访问池子
func DatabasePools(cfg *DbConfig, max int) *gorm.DB {
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
	return db
}
