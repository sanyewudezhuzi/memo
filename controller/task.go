package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanyewudezhuzi/memo/pkg/errcode"
	"github.com/sanyewudezhuzi/memo/pkg/util"
	"github.com/sanyewudezhuzi/memo/serializer"
	"github.com/sanyewudezhuzi/memo/service"
)

// 新建备忘录
func CreateTask(ctx *gin.Context) {
	// 获取 claim
	claim, ok := ctx.Get("claim")
	if !ok {
		ctx.JSON(http.StatusOK, serializer.Response{
			StatusCode: errcode.Failed_to_verify_identity,
			Error:      "Failed to verify identity.",
		})
	}

	// 创建服务
	var createTask service.CreateTaskService

	// 绑定参数
	if err := ctx.ShouldBind(&createTask); err != nil {
		ctx.JSON(http.StatusInternalServerError, serializer.Response{
			StatusCode: errcode.Parameter_transfer_error,
			Error:      "Parameter transfer error.",
		})
	} else {
		// 完成创建
		res := createTask.Create(claim.(*util.Claims).UID)
		ctx.JSON(http.StatusOK, res)
	}

}
