package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test-todolist/consts"
	"test-todolist/service"
	"test-todolist/types"
)

// CreateTaskHandler @Tags TASK
// @Summary 创建任务
// @Tags 任务模块
// @param title formData string false "任务名称"
// @param content formData string false "任务内容"
// @Success 200 {string} json{"code", "message"}
// @Router /task/createTask [post]
func CreateTaskHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.CreateTaskReq
		req.Content = context.PostForm("content")
		req.Title = context.PostForm("title")
		if err := context.ShouldBind(&req); err != nil {
			context.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}
		validateParam := service.GetTaskSrv()
		resp, err := validateParam.CreateTask(context.Request.Context(), &req)

		if err != nil {
			context.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		context.JSON(http.StatusOK, resp)
	}
}

// GetTaskListHandler @Tags TASK
// @Summary 查询任务列表
// @Tags 任务模块
// @param currentPage query string false "页码"
// @param pageSize query string false "页数"
// @Success 200 {string} json{"code", "message"}
// @Router /task/getTaskList [post]
func GetTaskListHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.ListTaskReq
		currentPage, _ := strconv.Atoi(context.Query("currentPage"))
		req.CurrentPage = currentPage
		pageSize, _ := strconv.Atoi(context.Query("pageSize"))
		req.PageSize = pageSize

		if req.PageSize == 0 {
			req.PageSize = consts.BasePageLimit
		}

		l := service.GetTaskSrv()
		resp, err := l.GetTaskList(context.Request.Context(), &req)
		if err != nil {
			context.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}
		context.JSON(http.StatusOK, resp)
	}
}

// UpdateTaskHandler @Tags TASK
// @Summary 编辑任务列表
// @Tags 任务模块
// @param title formData string false "标题"
// @param content formData string false "内容"
// @Success 200 {string} json{"code", "message"}
// @Router /task/updateTask [post]
func UpdateTaskHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.UpdateListTaskReq
		req.Title = context.PostForm("title")
		req.Content = context.PostForm("content")

		if err := context.ShouldBind(&req); err != nil {
			context.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}
		l := service.GetTaskSrv()
		resp, err := l.UpdateTask(context.Request.Context(), &req)
		if err != nil {
			context.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}
		context.JSON(http.StatusOK, resp)
	}
}

// DeleteTaskHandler @Tags TASK
// @Summary 删除任务列表
// @Tags 任务模块
// @param id formData string false "用户id"
// @Success 200 {string} json{"code", "message"}
// @Router /task/deleteTask [post]
func DeleteTaskHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.DeleteTaskReq

		if err := context.ShouldBind(&req); err != nil {
			context.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		l := service.GetTaskSrv()
		resp, err := l.DeleteTask(context.Request.Context(), &req)

		if err != nil {
			context.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		context.JSON(http.StatusOK, resp)

	}
}
