package utils

import "net/http"

type ErrorDto struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

// 错误

func NewError(code int, msg string) *ErrorDto {
	return &ErrorDto{
		ErrorCode: code,
		ErrorMsg:  msg,
	}
}

// 500 错误处理
func ServerError() *ErrorDto {
	return NewError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

// 404 错误
func NotFound() *ErrorDto {
	return NewError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

// 未知错误
func UnknownError(message string) *ErrorDto {
	return NewError(http.StatusForbidden, message)
}

// 参数错误
func ParameterError(message string) *ErrorDto {
	return NewError(http.StatusBadRequest, message)
}

// 授权错误
func AuthError(message string) *ErrorDto {
	return NewError(http.StatusBadRequest, message)
}

type SuccessMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewOkMsg(code int, msg string) SuccessMsg {
	var res SuccessMsg
	res.Code = code
	res.Msg = msg
	return res
}
func OkMsg(msg string) SuccessMsg {
	return NewOkMsg(http.StatusOK, msg)
}

type Success struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func NewSuccess(code int, msg string, data interface{}) Success {
	var res Success
	res.Result = data
	res.Code = code
	res.Msg = msg
	return res
}

func Ok(data interface{}) Success {
	return NewSuccess(http.StatusOK, "请求成功", data)
}

type Page struct {
	Msg       string      `json:"msg"`
	Count     int         `json:"count"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
	List      interface{} `json:"list"`
}
