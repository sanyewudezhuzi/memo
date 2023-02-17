package router

import (
	"github.com/sanyewudezhuzi/memo/controller"
	"github.com/sanyewudezhuzi/memo/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// 路由
func Router() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("thisisanewstore"))
	r.Use(sessions.Sessions("mysession", store))
	memo := r.Group("memo")
	{
		// 用户操作
		user := memo.Group("user")
		{
			user.POST("register", controller.UserRegister)
			user.POST("login", controller.UserLogin)
		}
		// 中间件
		memo.Use(middleware.JWT())
		// 备忘录操作
		task := memo.Group("task")
		{
			task.POST("create", controller.CreateTask)
			task.GET("show", controller.ShowTask)
			task.GET("list", controller.ListTask)
			task.PUT("update", controller.UpdateTask)
		}
	}
	return r
}
