package scope

import "github.com/jinzhu/gorm"

//筛查已删除记录
func DeletedRecords(db *gorm.DB) *gorm.DB {
	return db.Where("deleted_at is not null")
}

//筛查有效记录（未删除）（注意弥补分页时查询会所有记录Bug而用）
func LivedRecords(db *gorm.DB) *gorm.DB {
	return db.Where("deleted_at is null")
}
