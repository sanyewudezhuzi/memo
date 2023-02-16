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
