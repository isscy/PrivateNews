package logic

import (
	"serverGo/dao"
	"serverGo/foundation/jwt"
	"serverGo/model"
)

// Login 登录相关
func Login(userParam *model.LoginForm) (user *model.User, err error) {
	user = &model.User{
		Username: userParam.UserName,
		Password: userParam.Password,
		Keyword:  userParam.Keyword,
	}
	if err := dao.Login(user); err != nil {
		return nil, err
	}
	accessToken, refreshToken, err := jwt.GenToken(user.ID, user.Username)
	if err != nil {
		return
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	return
}

// UserInfo 获取用户信息
func UserInfo(uId uint64) (user *model.User, err error) {
	return dao.GetUserInfo(uId)
}
