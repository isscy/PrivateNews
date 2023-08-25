package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	ContextUserIdKey    = "UserId"
	ContextUserNameKey  = "UserName"
	ContextUserRolesKey = "UserRoles"
)

// GetCurrentUserID 获取当前用户
func GetCurrentUserID(c *gin.Context) (userID uint64, username string, err error) {
	_uId, ok := c.Get(ContextUserIdKey)
	if !ok {
		err = errors.New("无法获取用户id")
		return
	}
	_uName, ok := c.Get(ContextUserNameKey)
	if !ok {
		err = errors.New("无法获取用户名称")
		return
	}
	userID, ok = _uId.(uint64)
	if !ok {
		err = errors.New("userId格式错误")
		return
	}
	username, ok = _uName.(string)
	return
}

// GetPageInfo 分页参数
func GetPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	SizeStr := c.Query("size")
	var (
		page int64 // 第几页 页数
		size int64 // 每页几条数据
		err  error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(SizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
