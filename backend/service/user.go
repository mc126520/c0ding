package service

import (
	"c0ding-backend/model"
	"c0ding-backend/utils"
	"errors"
)

// 用户注册
func UserRegister(username string, password string) error {
	_, err := dao.GetUsername(username)
	if err != nil {
		return errors.New("用户名已存在")
	}
	user := &model.User{
		UserName:         username,
		Password:         password,
		Level:            0,
		Exp:              0,
		RecvLikes:        0,
		RecvSubscription: 0,
	}
	return dao.AddUser(user)
}

// 用户登录

func UserLogin(username, password string) (*model.User, error) {
	user, err := dao.GetUserName(username)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	pwd, err := dao.GetUserPassword(username)
	if err != nil {
		return nil, errors.New("无法查询注册密码")
	}
	if !utils.CheckPassword(password, pwd) {
		return nil, errors.New("密码不正确")
	}
	return user, nil
}
