package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"serverGo/foundation/resp"
	"serverGo/foundation/utils"
	"serverGo/logic"
	"serverGo/model"
	"strconv"
)

// 用户登录
func LoginController(c *gin.Context) {
	// 获取参数 及 参数验证
	var userParam *model.LoginForm
	if err := c.ShouldBindJSON(&userParam); err != nil {
		zap.L().Error("登录的请求参数错误！", zap.Error(err))
		errMsg := utils.GetValigMsg(err, userParam)
		fmt.Println("errMsg:  ", errMsg)
		resp.Fails(c, resp.InvalidParams, errMsg)
		return
	}
	// 用户校验
	user, err := logic.Login(userParam)
	if err != nil {
		zap.L().Error("用户登录失败：", zap.String("username", userParam.UserName), zap.Error(err))
		resp.Fails(c, resp.BizCheckFail, err.Error())
		return
	}
	// 返回响应
	resp.Succ(c, user)
}

// 获取用户信息
func UserInfoController(c *gin.Context) {
	// 获取参数 及 参数验证
	userIdStr, exists := c.Get(utils.ContextUserIdKey)
	if !exists {
		resp.Fails(c, resp.InvalidToken, "无法获取用户信息")
	}
	uId := userIdStr.(uint64)
	user, err := logic.UserInfo(uId)
	if err != nil {
		zap.L().Error("获取用户信息失败：", zap.String("userId:  ", userIdStr.(string)), zap.Error(err))
		resp.Fails(c, resp.BizCheckFail, err.Error())
		return
	}
	// 返回响应
	resp.Succ(c, user)
}

func PanicController(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			err.(error).Error()
			c.String(http.StatusOK, err.(error).Error())
		}
	}()
	params, _ := strconv.Atoi(c.Query("num"))
	i := 10 / params
	c.String(http.StatusOK, strconv.Itoa(i))
}

// 连接测试
func TestController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "ping"})
}
