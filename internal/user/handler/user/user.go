package handler

import (
	"jet-ci/internal/user/domain/userinfo"
	"jet-ci/internal/user/repository"
)

type IUserHandler interface {
	CreateUser(*userinfo.InputCreateUser) error
}

type UserHandler struct {
	UserRepository repository.IUserRepository
}

func NewUserHandler() IUserHandler {
	h := &UserHandler{
		//这玩意也可以传过来
		UserRepository: repository.NewUserRepository(),
	}
	return h
}

func (h *UserHandler) CreateUser(input *userinfo.InputCreateUser) error {
	model := &userinfo.Account{
		UserName:  input.UserName,
		Email:     input.Email,
		Name:      input.Name,
		CellPhone: input.CellPhone,
	}
	pwd, err := model.GeneratePassword(input.Password)
	if err != nil {
		return err
	}
	model.Password = string(pwd)
	err = h.UserRepository.CreateUser(model)
	return err
}
