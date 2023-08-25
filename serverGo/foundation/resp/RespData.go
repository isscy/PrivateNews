package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespData struct {
	Code    BizCode     `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // omitempty当data为空时,不展示这个字段
}

func Fail(ctx *gin.Context, c BizCode) {
	resultData := &RespData{
		Code:    c,
		Message: c.Msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, resultData)
}

func Fails(ctx *gin.Context, c BizCode, msg string) {
	resultData := &RespData{
		Code:    c,
		Message: msg,
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, resultData)
}

func Succ(ctx *gin.Context, data interface{}) {
	rd := &RespData{
		Code:    Success,
		Message: Success.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}
