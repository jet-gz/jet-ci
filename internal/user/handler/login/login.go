package handler

import (
	"jet-ci/internal/user/domain/userinfo"
	"jet-ci/internal/user/repository"
)

type ILoginHandler interface {
	Login(userinfo.InputLogin) (userinfo.OutputLogin, error)
}

type LoginHandler struct {
	UserRepository repository.IUserRepository
}

func NewLoginHandler() ILoginHandler {
	h := &LoginHandler{
		//这玩意也可以传过来
		UserRepository: repository.NewUserRepository(),
	}
	return h
}

func (h *LoginHandler) Login(input userinfo.InputLogin) (user userinfo.OutputLogin, err error) {
	u, err := h.UserRepository.FindUserByName(input.UserName)
	if err != nil {
		return user, err
	}
	isOk, err := u.CheckPwd(input.Password)
	if !isOk {
		return user, err
	}
	user.Avatar = "111111111111111111111111"

	//生成Token
	return user, err
}
