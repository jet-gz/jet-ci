package userinfo

import (
	"errors"
	"fmt"
	"jet-ci/internal/pkg/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserName  string `gorm:"size:64;comment:用户名"`
	Name      string `gorm:"size:64;comment:姓名"`
	Email     string `gorm:"size:64;comment:邮箱"`
	CellPhone string `gorm:"size:64;comment:手机"`
	Password  string `gorm:"size:64;comment:密码"`
	Disabled  bool   `gorm:"size:64;comment:禁用"`
	//Roles     []role.Role
}

func (a *Account) New() Account {
	return Account{
		Name:  "jet",
		Email: "jet-gg@sds.com",
	}
}

func (a *Account) CheckPwd(password string) (isOk bool, err error) {

	if err = bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password)); err != nil {
		return false, errors.New("密码比对错误")
	}
	return true, nil
}

func (a *Account) GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

type Token struct {
}

type InputLogin struct {
	utils.IRequestDto
	UserName string `json:"username" validate:"required" label:"用户名或邮箱"`
	Password string `json:"password" validate:"required" label:"密码"`
	Code     string `json:"code"`
	Key      string `json:"key"`
}

func (input *InputLogin) Validate() error {

	fmt.Println("dto验证")
	return nil
}

type OutputLogin struct {
	Avatar string
	Token  Token
}

type InputCreateUser struct {
	utils.IRequestDto
	UserName  string `json:"username"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CellPhone string `json:"cellPhone"`
	Password  string `json:"password"`
}

func (input *InputCreateUser) Validate() error {

	fmt.Println("dto验证")
	return nil
}
