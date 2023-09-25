package api

import (
	"jet-ci/internal/pkg/utils"
	"jet-ci/internal/user/domain/userinfo"
	handler "jet-ci/internal/user/handler/login"

	"github.com/gin-gonic/gin"
)

type LoginApi struct {
	utils.Api
	LoginHandler handler.ILoginHandler
}

func NewLoginApi() LoginApi {
	return LoginApi{
		LoginHandler: handler.NewLoginHandler(),
	}
}

func LoginRouting(r *gin.RouterGroup) {
	api := NewLoginApi()
	r.POST("/login", api.Login)

}

func (api *LoginApi) Login(ctx *gin.Context) {
	var input userinfo.InputLogin
	api.MakeContext(ctx, &input)
	out, err := api.LoginHandler.Login(input)
	if err != nil {
		api.Error(err.Error())
	}
	api.OK(out)
}
