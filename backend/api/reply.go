package api

import (
	"c0ding-backend/response"
	"c0ding-backend/service"

	"github.com/gin-gonic/gin"
)

type Reply struct {
	UserName string
	Reply    string
}

// 回复
func UploadReply(c *gin.Context) {
	var upreply Reply
	if err := c.ShouldBindJSON(&upreply); err != nil {
		response.FailBadRequest(c, 400, "")
		return
	}
	if err := service.UploadReply(); err != nil {
		response.FailBadRequest(c, 400, "")
		return
	}
	response.Success(c, 200, "")
}

// 删除回复
func DeleteReply(c *gin.Context) {
	var Index uint
	if err := service.DeleteReply(Index); err != nil {
		response.FailBadRequest(c, 400, "")
		return
	}
	response.Success(c, 200, "")
}
