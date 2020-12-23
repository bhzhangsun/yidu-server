package utils

const (
	RESPCODE_SUCCESS      = 0
	RESPCODE_ERROR_SERVER = 5000
	RESPCODE_ERROR_ROUTE  = 4000
	RESPCODE_ERROR_PARAM  = 4001
)

const (
	RESPMSG_SUCCESS      = "OK"
	RESPMSG_ERROR_SERVER = "服务器内部错误"
	RESPMSG_ERROR_PARAM  = "传入参数错误"
	RESPMSG_ERROR_ROUTE  = "资源未找到"
)

type Response struct {
	Code    int         `json: "code"`
	Data    interface{} `json: "data"`
	Message string      `json: "message"`
}
