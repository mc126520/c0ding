package service

// 回复
func UploadReply() error {
	return dao.UploadReply()
}

func DeleteReply(Index uint) error {
	return dao.DeleteReply(Index)
}
