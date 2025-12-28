package api

import (
	"c0ding-backend/response"
	"c0ding-backend/service"

	"github.com/gin-gonic/gin"
)

// 点赞
func UploadLikes(c *gin.Context) {
	if err := service.UploadLikes(); err != nil {
		response.FailBadRequest(c, 400, "")
		return
	}
	response.Success(c, 200, "")
}

// 取消点赞
func CancelLikes(c *gin.Context) {
	if err := service.CancelLikes(); err != nil {
		response.FailBadRequest(c, 400, "")
		return
	}
	response.Success(c, 200, "")
}
