package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test-todolist/service"
	"test-todolist/types"
)

// RegisterUser
// @Summary 新增用户
// @Tags 用户模块
// @param username formData string false "用户名"
// @param password formData string false "密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/register [post]
func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserServiceReq
		req.UserName = c.PostForm("username")
		req.Password = c.PostForm("password")
		fmt.Println(c.PostForm("username"), c.PostForm("password"), req.UserName, req.Password, "111111111")

		// req结合ShouldBind对数据做校验
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		validateParam := service.GetUserSrv()
		resp, err := validateParam.Register(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

// UserLoginHandler
// @Summary 登录
// @Tags 用户模块
// @param username formData string false "用户名"
// @param password formData string false "密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/login [post]
func UserLoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserServiceReq
		req.UserName = c.PostForm("username")
		req.Password = c.PostForm("password")
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		validateParam := service.GetUserSrv()
		resp, err := validateParam.Login(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, resp)

	}
}
