package resp

type BizCode int64

const (
	Success           BizCode = 1000
	InvalidParams     BizCode = 1001
	InvalidToken      BizCode = 1002
	InvalidAuthFormat BizCode = 1003
	BizCheckFail      BizCode = 1004
	InvalidPath       BizCode = 1005
	ServerUnKnowErr   BizCode = 2001
	ServerBusy        BizCode = 2002
)

var msgFlags = map[BizCode]string{
	Success:           "success",
	InvalidParams:     "请求参数错误",
	InvalidToken:      "无效的Token",
	InvalidAuthFormat: "认证格式有误",
	BizCheckFail:      "业务校验不通过",
	InvalidPath:       "404：无效的路径",
	ServerUnKnowErr:   "服务未知错误",
	ServerBusy:        "服务繁忙",
}

// (c BizCode) 表示该func从属于BizCode这个结构体 并可以通过c调用结构体的方法和参数
func (c BizCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[ServerUnKnowErr]
}
