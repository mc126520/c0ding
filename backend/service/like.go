package service

import "c0ding-backend/model"

// 点赞
func UploadLikes() error {
	var LikeSum uint

	like := &model.Post{
		Likes: e,
	}
	return dao.UploadLikes(like)
}

// 取消点赞
func CancelLikes() error {
	return dao.CancelLikes(like)
}
