package middleware

import (
	"github.com/gin-gonic/gin"
	"test-todolist/pkg/ctl"
	"test-todolist/pkg/e"
	"test-todolist/pkg/utils"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		code = e.SUCCESS

		token := context.GetHeader("Authorization")

		if token == "" {
			code = e.ErrorCheckTokenNotFound
			context.JSON(e.InvalidParams, gin.H{
				"msg":    e.GetMessage(code),
				"status": code,
				"data":   "缺少token",
			})
			context.Abort()
			return
		}

		authToken, err := utils.ParseToken(token)

		if err != nil {
			code = e.ErrorCheckTokenFail
		} else if time.Now().Unix() > authToken.ExpiresAt {
			code = e.ErrorCheckTokenTimeout
		}

		if code != e.SUCCESS {
			context.JSON(e.InvalidParams, gin.H{
				"status": code,
				"data":   "token可能过期了，请重新登录",
				"msg":    e.GetMessage(code),
			})
			context.Abort()
			return
		}

		context.Request = context.Request.WithContext(ctl.NewContext(context.Request.Context(), &ctl.UserInfo{
			Id: authToken.Id,
		}))

		context.Next()

	}
}
