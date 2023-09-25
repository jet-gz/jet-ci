package db

import (
	"jet-ci/internal/user/domain/userinfo"
)

func Start() {
	// 配置文件
	Load()
	mysqlOpen()
	//migrate()

}

func migrate() {
	ORM.AutoMigrate(
		new(userinfo.Account),
	)
}
