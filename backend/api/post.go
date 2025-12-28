package api

import (
	"c0ding-backend/response"
	"c0ding-backend/service"

	"github.com/gin-gonic/gin"
)

type UploadP struct {
	AuthorName string
	Content    string
	PostID     uint
}

// 上传Post
func UploadPost(c *gin.Context) {
	var uploadpost UploadP
	if err := c.ShouldBindJSON(&uploadpost); err != nil {
		response.FailBadRequest(c, 400, "")
		return
	}
	if err := service.UploadPost(uploadpost.AuthorName, uploadpost.Content); err != nil {
		response.FailBadRequest(c, 400, "")
		return
	}
	response.Success(c, 200, "")
}

// 删除Post
func DeletePost(c *gin.Context) {
	var delpost UploadP
	if err := c.ShouldBindJSON(&delpost); err != nil {
		response.FailBadRequest(c, 400, "")
		return
	}
	if err := service.DeletePost(delpost.PostID); err != nil {
		response.FailBadRequest(c, 400, "")
		return
	}
	response.Success(c, 200, "")
}
