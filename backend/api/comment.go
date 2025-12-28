package api

import (
	"c0ding-backend/response"
	"c0ding-backend/service"

	"github.com/gin-gonic/gin"
)

type UploadC struct {
	UserName string
	Comment  string
}

// 上传评论
func UploadComment(c *gin.Context) {
	var uploadcomment UploadC
	if err := c.ShouldBindJSON(&uploadcomment); err != nil {
		response.FailBadRequest(c, 400, "")
		return
	}
	if err := service.UploadComment(uploadcomment.UserName, uploadcomment.Comment); err != nil {
		response.FailBadRequest(c, 400, "")
		return
	}
}

// 删除评论
func DeleteComment(c *gin.Context) {
	var Index uint
	if err := service.DeleteComment(Index); err != nil {
		response.FailBadRequest(c, 400, "")
		return
	}
	response.Success(c, 200, "")
}
