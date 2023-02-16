package dao

import "github.com/sanyewudezhuzi/memo/model"

func CreateTask(task *model.Task) error {
	return model.DB.Model(&model.Task{}).Create(task).Error
}
