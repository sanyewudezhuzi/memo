package controller

// 接口层

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

// 展示一条备忘录
func ShowTask(ctx *gin.Context) {
	// 获取 claim
	claim, ok := ctx.Get("claim")
	if !ok {
		ctx.JSON(http.StatusOK, serializer.Response{
			StatusCode: errcode.Failed_to_verify_identity,
			Error:      "Failed to verify identity.",
		})
	}

	// 创建服务
	var showTask service.ShowTaskService

	// 绑定参数
	if err := ctx.ShouldBind(&showTask); err != nil {
		ctx.JSON(http.StatusInternalServerError, serializer.Response{
			StatusCode: errcode.Parameter_transfer_error,
			Error:      "Parameter transfer error.",
		})
	} else {
		// 完成创建
		res := showTask.Show(claim.(*util.Claims).UID)
		ctx.JSON(http.StatusOK, res)
	}
}

// 展示所有备忘录
func ListTask(ctx *gin.Context) {
	// 获取 claim
	claim, ok := ctx.Get("claim")
	if !ok {
		ctx.JSON(http.StatusOK, serializer.Response{
			StatusCode: errcode.Failed_to_verify_identity,
			Error:      "Failed to verify identity.",
		})
	}

	// 创建服务
	var listTask service.ListTaskService

	// 绑定参数
	if err := ctx.ShouldBind(&listTask); err != nil {
		ctx.JSON(http.StatusInternalServerError, serializer.Response{
			StatusCode: errcode.Parameter_transfer_error,
			Error:      "Parameter transfer error.",
		})
	} else {
		res := listTask.List(claim.(*util.Claims).UID)
		ctx.JSON(http.StatusOK, res)
	}
}

// 更新备忘录
func UpdateTask(ctx *gin.Context) {
	claim, ok := ctx.Get("claim")
	if !ok {
		ctx.JSON(http.StatusOK, serializer.Response{
			StatusCode: errcode.Failed_to_verify_identity,
			Error:      "Failed to verify identity.",
		})
	}
	var updateTask service.UpdateTaskService
	if err := ctx.ShouldBind(&updateTask); err != nil {
		ctx.JSON(http.StatusInternalServerError, serializer.Response{
			StatusCode: errcode.Parameter_transfer_error,
			Error:      "Parameter transfer error.",
		})
	} else {
		res := updateTask.Update(claim.(*util.Claims).UID)
		ctx.JSON(http.StatusOK, res)
	}
}

// 删除备忘录
func DeleteTask(ctx *gin.Context) {
	claim, ok := ctx.Get("claim")
	if !ok {
		ctx.JSON(http.StatusOK, serializer.Response{
			StatusCode: errcode.Failed_to_verify_identity,
			Error:      "Failed to verify identity.",
		})
	}
	var deleteTask service.DeleteTaskService
	if err := ctx.ShouldBind(&deleteTask); err != nil {
		ctx.JSON(http.StatusInternalServerError, serializer.Response{
			StatusCode: errcode.Parameter_transfer_error,
			Error:      "Parameter transfer error.",
		})
	} else {
		res := deleteTask.Delete(claim.(*util.Claims).UID)
		ctx.JSON(http.StatusOK, res)
	}
}
