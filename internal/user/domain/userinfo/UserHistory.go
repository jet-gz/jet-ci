package userinfo

import "gorm.io/gorm"

type UserHistory struct {
	gorm.Model
	UserId string
	Event  int
	Desc   string
}
