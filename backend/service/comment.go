package service

import "c0ding-backend/model"

// 上传评论
func UploadComment(username string, comment string) error {
	post := &model.Comment{
		UserName: username,
		Content:  comment,
	}
	return dao.UploadComment(post)
}

// 删除评论
func DeleteComment(index uint) error {
	return dao.DeleteComment(index)
}
