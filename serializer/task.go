package serializer

// 序列化 task 模块

import "github.com/sanyewudezhuzi/memo/model"

type Task struct {
	TID       uint   `json:"tid"`
	CreateAt  int64  `json:"create_at"`
	UpdateAt  int64  `json:"update_at"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Status    string `json:"status"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

// 序列化 task
func BuildTask(task model.Task) Task {
	return Task{
		TID:       task.ID,
		CreateAt:  task.CreatedAt.Unix(),
		UpdateAt:  task.UpdatedAt.Unix(),
		Title:     task.Title,
		Content:   task.Content,
		Status:    task.Status,
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
}

// 序列化 tasks
func BuildTasks(tasks []model.Task) []Task {
	var task []Task = make([]Task, len(tasks))
	for k, v := range tasks {
		task[k] = Task{
			TID:       v.ID,
			CreateAt:  v.CreatedAt.Unix(),
			UpdateAt:  v.UpdatedAt.Unix(),
			Title:     v.Title,
			Content:   v.Content,
			Status:    v.Status,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		}
	}
	return task
}
