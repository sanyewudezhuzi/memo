package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Account        string `gorm:"unique"`
	PassWordDigest string
	UserName       string
}
