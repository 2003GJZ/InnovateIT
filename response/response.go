package response

//在此规定返回体结构

type Response struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 返回信息
	Data    interface{} `json:"data"`    // 返回数据
}

// NewResponse 创建一个新的返回体
func NewResponse(code int, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
