package model

import (
	"time"
)

type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AllTimestamps struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `sql:"index"`
}
