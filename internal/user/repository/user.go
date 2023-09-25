package repository

import (
	"jet-ci/internal/pkg/db"
	"jet-ci/internal/user/domain/userinfo"
)

type IUserRepository interface {
	// //初始化数据表
	// InitTable() error
	// //根据用户名称查找用户信息
	// FindUserByName(string) (*model.User,error)
	//根据用户ID查找用户信息
	FindUserByID(int64) (*userinfo.Account, error)
	// //创建用户
	CreateUser(*userinfo.Account) error

	//用户名或邮箱查询
	FindUserByName(name string) (account *userinfo.Account, err error)
}

type UserRepository struct {
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (u *UserRepository) FindUserByID(id int64) (account *userinfo.Account, err error) {
	account = &userinfo.Account{}
	db.ORM.First(account, id)
	return account, err
}
func (u *UserRepository) FindUserByName(name string) (account *userinfo.Account, err error) {
	account = &userinfo.Account{}
	return account, db.ORM.Where(&userinfo.Account{UserName: name}).Or(&userinfo.Account{Email: name}).Find(account).Error
}
func (u *UserRepository) CreateUser(model *userinfo.Account) error {
	return db.ORM.Create(model).Error
}
