package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Api struct {
	Context *gin.Context
	Errors  error
}

func (e *Api) MakeContext(ctx *gin.Context, input IRequestDto) *Api {
	e.Context = ctx
	if err := ctx.ShouldBindJSON(&input); err != nil {
		fmt.Println("---->", err)
		e.Error("请求异常")
		ctx.Abort()
	}
	err := input.Validate()
	if err != nil {
		e.Error("请求异常")
		ctx.Abort()
	}
	return e
}

func (e *Api) ValidatorError(c *gin.Context, err error) {

}

func (api *Api) Page(result interface{}, count int, pageIndex int, pageSize int, msg string) {
	res := Page{
		List:      result,
		PageIndex: pageIndex,
		PageSize:  pageSize,
		Msg:       msg,
		Count:     count,
	}
	api.Context.AbortWithStatusJSON(http.StatusOK, res)

}

func (api *Api) Error(msg string) {
	var res ErrorDto
	res.ErrorCode = -1
	res.ErrorMsg = msg
	api.Context.AbortWithStatusJSON(http.StatusForbidden, res)
}

func (api *Api) BizError(code int, msg string) {
	var res ErrorDto
	res.ErrorCode = -1
	res.ErrorMsg = msg
	api.Context.AbortWithStatusJSON(http.StatusForbidden, res)
}

// error 可以搞个自定义的
func (api *Api) Exception(code int, err error, msg string) {

}

func (api *Api) OkMsg(msg string) {
	var res SuccessMsg
	res.Code = http.StatusOK
	res.Msg = msg
	api.Context.AbortWithStatusJSON(http.StatusForbidden, res)
}

func (api *Api) OK(data interface{}) {
	var res Success
	res.Result = data
	res.Code = http.StatusOK
	res.Msg = "请求成功"
	api.Context.AbortWithStatusJSON(http.StatusOK, res)
}
