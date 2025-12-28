package api

import (
	"c0ding-backend/response"
	"c0ding-backend/service"
	"c0ding-backend/utils"

	"github.com/gin-gonic/gin"
)

type UserRegRes struct {
	UserName string
	Password string
}
type UserLoginRes struct {
	UserName string
	Password string
}

// 用户注册
func UserRegister(c *gin.Context) {
	var useraddreq UserRegRes
	if err := c.ShouldBindJSON(&useraddreq); err != nil {
		response.FailBadRequest(c, 400, "")
		return
	}
	hashpassword, err := utils.HashPassword(useraddreq.Password)
	if err != nil {
		response.FailBadRequest(c, 400, "")
		return
	}
	useraddreq.Password = hashpassword
	if err := service.UserRegister(useraddreq.UserName, useraddreq.Password); err != nil {
		response.FailBadRequest(c, 400, "")
		return
	}
	response.Success(c, 200, "")
}

// 用户登录
func UserLogin(c *gin.Context) {
	var userloginres UserLoginRes
	if err := c.ShouldBindJSON(&userloginres); err != nil {
		response.FailBadRequest(c, 400, "")
		return
	}
	_, err := service.UserLogin(userloginres.UserName, userloginres.Password)
	if err != nil {
		response.FailBadRequest(c, 400, "")
		return
	}

	response.Success(c, 200, "登陆成功!")
}
