package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"test-todolist/api"
	docs "test-todolist/docs"
	"test-todolist/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 创建基于加密的 Cookie 存储实例
	store := cookie.NewStore([]byte("todo-secret"))

	// 使用 Cookie 存储创建会话中间件
	r.Use(sessions.Sessions("my-session", store))

	r.Use(middleware.Cors())

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.POST("/user/register", api.RegisterUser())
		v1.POST("/user/login", api.UserLoginHandler())
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("/task/createTask", api.CreateTaskHandler())
			authed.GET("task/getTaskList", api.GetTaskListHandler())
			authed.POST("task/updateTask", api.UpdateTaskHandler())
			authed.POST("/task/deleteTask", api.DeleteTaskHandler())
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
