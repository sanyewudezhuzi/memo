package service

import (
	"fmt"
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

func (s *CreateTaskService) Create(uid uint) serializer.Response {
	status_code := 0
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
	fmt.Println(task)
	err := dao.CreateTask(&task)
	if err != nil {
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
