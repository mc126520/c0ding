package service

import (
	"c0ding-backend/model"
	"time"
)

// 上传post
func UploadPost(author string, content string) error {
	post := &model.Post{
		AuthorName: author,
		Content:    content,
		Likes:      0,
		PostTime:   time.Now(),
		Views:      0,
	}
	return dao.UploadPost(post)
}

// 修改post
func UpdatePost(author string, content string) error {
	post := &model.Post{
		Content:    content,
		UpdateTime: time.Now(),
	}
	return dao.UpdatePost(post)
}

// 删除post
func DeletePost(postId uint) error {
	return dao.DeletePost(postId)
}
