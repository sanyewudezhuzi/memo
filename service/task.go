package service

import (
	"time"

	"github.com/sanyewudezhuzi/memo/dao"
	"github.com/sanyewudezhuzi/memo/model"
	"github.com/sanyewudezhuzi/memo/pkg/errcode"
	"github.com/sanyewudezhuzi/memo/serializer"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content"`
	Status  string `json:"status" form:"status"`
}

type ShowTaskService struct {
	Title string `json:"title" form:"title" binding:"required"`
}

type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

type UpdateTaskService struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content"`
	Status  string `json:"status" form:"status"`
}

type DeleteTaskService struct {
	Title string `json:"title" form:"title" binding:"required"`
}

func (s *CreateTaskService) Create(uid uint) serializer.Response {
	status_code := 0
	var count int64
	model.DB.Model(&model.Task{}).Where("title = ? and user_id = ?", s.Title, uid).Count(&count)
	if count != 0 {
		status_code = errcode.The_title_has_been_used
		return serializer.Response{
			StatusCode: status_code,
			Error:      "The title has been used.",
		}
	}
	var endTime int64 = 0
	if s.Status == "" {
		s.Status = "Continue"
	} else if s.Status == "Done" {
		endTime = time.Now().Unix()
	}
	var task = model.Task{
		User_ID:   uid,
		Title:     s.Title,
		Content:   s.Content,
		Status:    s.Status,
		StartTime: time.Now().Unix(),
		EndTime:   endTime,
	}
	if err := dao.CreateTask(&task); err != nil {
		status_code = errcode.Create_task_error
		return serializer.Response{
			StatusCode: status_code,
			Error:      "Create task error.",
		}
	}
	return serializer.Response{
		StatusCode: errcode.OK,
		Msg:        "Create task successed.",
	}
}

func (s *ShowTaskService) Show(uid uint) serializer.Response {
	status_code := 0
	var task model.Task
	// 此处可能需要防止 sql 注入
	model.DB.Model(&model.Task{}).Where("user_id = ? and title = ?", uid, s.Title).First(&task)

	// 判断 task 是否存在
	if task.ID == 0 {
		status_code = errcode.No_title_found
		return serializer.Response{
			StatusCode: status_code,
			Error:      "No title found.",
		}
	}

	// 返回响应
	return serializer.Response{
		StatusCode: errcode.OK,
		Data:       serializer.BuildTask(task),
		Msg:        "Show task successed.",
	}
}

func (s *ListTaskService) List(uid uint) serializer.Response {
	var tasks []model.Task
	var count int64
	if s.PageSize == 0 {
		s.PageSize = 5
	}
	if s.PageNum == 0 {
		s.PageNum = 1
	}
	// 此处可能需要防止 sql 注入
	model.DB.Model(&model.Task{}).Where("user_id = ?", uid).
		Limit(s.PageSize).Offset((s.PageNum - 1) * s.PageSize).
		Find(&tasks).Count(&count)
	return serializer.Response{
		StatusCode: errcode.OK,
		Data: serializer.ListData{
			List:  serializer.BuildTasks(tasks),
			Total: int(count),
		},
		Msg: "List task successed.",
	}
}

func (s *UpdateTaskService) Update(uid uint) serializer.Response {
	// 标题不能更改
	status_code := 0
	var task model.Task
	model.DB.Model(&model.Task{}).Where("title = ? and user_id = ?", s.Title, uid).First(&task)
	if task.Title == "" {
		status_code = errcode.No_title_found
		return serializer.Response{
			StatusCode: status_code,
			Error:      "No title found.",
		}
	}
	if s.Content != "" {
		task.Content = s.Content
	}
	if s.Status != "" {
		task.Status = s.Status
	}
	if err := dao.UpdateTask(&task); err != nil {
		status_code = errcode.Update_task_error
		return serializer.Response{
			StatusCode: status_code,
			Error:      "Update task error.",
		}
	}
	return serializer.Response{
		StatusCode: errcode.OK,
		Data:       serializer.BuildTask(task),
		Msg:        "Update task successed.",
	}
}

func (s *DeleteTaskService) Delete(uid uint) serializer.Response {
	status_code := 0
	var task model.Task
	model.DB.Model(&model.Task{}).Where("user_id = ? and title = ?", uid, s.Title).First(&task)
	if task.Title == "" {
		status_code = errcode.No_title_found
		return serializer.Response{
			StatusCode: status_code,
			Error:      "No title found.",
		}
	}
	if err := dao.DeleteTask(&task); err != nil {
		status_code = errcode.Delete_task_error
		return serializer.Response{
			StatusCode: status_code,
			Error:      "Delete task error.",
		}
	}
	return serializer.Response{
		StatusCode: errcode.OK,
		Msg:        "Delete task successed.",
	}
}
