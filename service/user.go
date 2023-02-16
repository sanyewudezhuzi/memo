package service

import (
	"github.com/sanyewudezhuzi/memo/dao"
	"github.com/sanyewudezhuzi/memo/model"
	"github.com/sanyewudezhuzi/memo/pkg/errcode"
	"github.com/sanyewudezhuzi/memo/serializer"
)

type UserRegisterService struct {
	Account  string `form:"account" json:"account" binding:"required,min=1,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=4,max=25"`
	UserName string `form:"user_name" json:"user_name"`
}

func (s *UserRegisterService) Register() serializer.Response {
	// 创建模型
	status_code := 0
	var user model.User

	// 验证用户是否已存在
	model.DB.Model(&model.User{}).Where("account = ?", s.Account).First(&user)
	if user.ID != 0 {
		status_code = errcode.User_already_exists
		return serializer.Response{
			StatusCode: status_code,
			Error:      "User already exists.",
		}
	}

	// 加密密码
	if err := user.Encrypt(s.Password); err != nil {
		status_code = errcode.Bcrypt_error
		return serializer.Response{
			StatusCode: status_code,
			Error:      err.Error(),
		}
	}

	// 封装 user
	user.Account = s.Account
	user.UserName = user.Generate(s.UserName)

	// 数据持久化
	if err := dao.CreateUser(&user); err != nil {
		status_code = errcode.Create_user_error
		return serializer.Response{
			StatusCode: status_code,
			Error:      err.Error(),
		}
	}

	// 返回响应
	return serializer.Response{
		StatusCode: errcode.OK,
		Msg:        "Successful registration.",
	}
}
