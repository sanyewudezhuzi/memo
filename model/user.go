package model

// user 模块

import (
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account        string `gorm:"unique"`
	PassWordDigest string
	UserName       string
}

// 加密
func (u *User) Encrypt(pwd string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		return err
	}
	u.PassWordDigest = string(hash)
	return nil
}

// 检验
func (u *User) Check(pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PassWordDigest), []byte(pwd))
	return err == nil
}

// 生成用户名
func (u *User) Generate(username string) string {
	rand.Seed(time.Now().Unix())
	if username == "" {
		id := rand.Intn(10000)
		for id < 1000 || id > 10000 {
			if id < 1000 {
				id += rand.Intn(10000)
			} else {
				id -= rand.Intn(10000)
			}
		}
		username = "用户" + strconv.Itoa(rand.Intn(10000))
	}
	return username
}
