package dao

import "github.com/sanyewudezhuzi/memo/model"

func CreateTask(task *model.Task) error {
	return model.DB.Model(&model.Task{}).Create(task).Error
}

func UpdateTask(task *model.Task) error {
	return model.DB.Save(task).Error
}

func DeleteTask(task *model.Task) error {
	return model.DB.Delete(task).Error
}
