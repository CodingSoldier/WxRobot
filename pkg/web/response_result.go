package web

type ResponseResult struct {
	Code    int         `json:"code"`    //提示代码
	Message string      `json:"message"` //提示信息
	Data    interface{} `json:"data"`    //数据
}

func Success() *ResponseResult {
	return &ResponseResult{0, "success", nil}
}

func SuccessData(data interface{}) *ResponseResult {
	return &ResponseResult{0, "success", data}
}

func Fail(message string) *ResponseResult {
	return &ResponseResult{-1, message, nil}
}
