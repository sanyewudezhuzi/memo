package controller

import (
	"net/http"

	"github.com/sanyewudezhuzi/memo/pkg/errcode"
	"github.com/sanyewudezhuzi/memo/serializer"
	"github.com/sanyewudezhuzi/memo/service"

	"github.com/gin-gonic/gin"
)

// 用户注册
func UserRegister(ctx *gin.Context) {
	// 创建服务
	var userRegister service.UserRegisterService
	// 绑定参数
	if err := ctx.ShouldBind(&userRegister); err != nil {
		ctx.JSON(http.StatusInternalServerError, serializer.Response{
			StatusCode: errcode.Invalid_pass_parameter,
			Error:      "Invalid pass parameter.",
		})
	} else {
		// 完成注册
		res := userRegister.Register()
		ctx.JSON(http.StatusOK, res)
	}
}

// 用户登录
func UserLogin(ctx *gin.Context) {
	var userLogin service.UserLoginService
	if err := ctx.ShouldBind(&userLogin); err != nil {
		ctx.JSON(http.StatusInternalServerError, serializer.Response{
			StatusCode: errcode.Invalid_pass_parameter,
			Error:      "Invalid pass parameter.",
		})
	} else {
		res := userLogin.Login()
		ctx.JSON(http.StatusOK, res)
	}
}
