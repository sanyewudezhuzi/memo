package dao

import "github.com/sanyewudezhuzi/memo/model"

func CreateUser(user *model.User) error {
	return model.DB.Model(&model.User{}).Create(user).Error
}
