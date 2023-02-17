package middleware

// 中间件

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sanyewudezhuzi/memo/pkg/errcode"
	"github.com/sanyewudezhuzi/memo/pkg/util"
)

// 解析 token
func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status_code := 0
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			status_code = errcode.Failed_to_get_request_header
		} else {
			claim, err := util.ParseToken(tokenString)
			if err != nil {
				status_code = errcode.Failed_to_load_token
			} else if time.Now().Unix() > claim.ExpiresAt {
				status_code = errcode.Token_has_expired
			} else {
				ctx.Set("claim", claim)
			}
		}
		if status_code != 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": status_code,
				"msg":         "Failed to parse token.",
			})
			ctx.Abort()
			return
		}
	}
}
