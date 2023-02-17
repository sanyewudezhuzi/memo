package model

// task 模块

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	User_ID   uint   `gorm:"not null"`
	Title     string `gorm:"index; not null"`
	Content   string `gorm:"type:longtext"`
	Status    string `gorm:"default:'Continue'"` // Done or Continue
	StartTime int64
	EndTime   int64
}
