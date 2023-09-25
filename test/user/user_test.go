package user

import (
	"fmt"
	"jet-ci/internal/user/domain/userinfo"
	"testing"
)

func TestCheckPwd(t *testing.T) {
	u := &userinfo.Account{Password: "$2a$10$I0I.Pq5/U9zrkFSHmYtq7Oz45nC2zOb3.t86GO5PdoClw6w7iLe5S"} //123456
	// p, _ := u.GeneratePassword("123456")
	// newPwd := string(p)
	// fmt.Println("新密码", newPwd)

	if isok, err := u.CheckPwd("123456"); isok && err == nil {
		fmt.Println("密码验证通过...")
	}
}
