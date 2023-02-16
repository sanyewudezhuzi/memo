package router

import (
	"github.com/sanyewudezhuzi/memo/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

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
		}
	}
	return r
}
