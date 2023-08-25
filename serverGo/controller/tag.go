package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"serverGo/foundation/resp"
	"serverGo/logic"
	"serverGo/model"
)

// 获取内置tagLsit
func BuiltInTagController(c *gin.Context) {
	// 获取参数 及 参数验证
	var tagParam *model.BuiltInTagForm = new(model.BuiltInTagForm)

	tagList, err := logic.GetBuiltInTagList(tagParam)
	if err != nil {
		zap.L().Error("查询内置Tag的请求参数错误：", zap.Error(err))
		resp.Fails(c, resp.BizCheckFail, err.Error())
		return
	}
	// 返回响应
	resp.Succ(c, tagList)
}
