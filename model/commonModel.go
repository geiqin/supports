package model

import (
	"time"
)

//基本Model
type BaseModel struct {
	Id        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type IntModel struct {
	Id        int32 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}


