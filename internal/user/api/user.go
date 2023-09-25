package api

import (
	"jet-ci/internal/pkg/utils"
	"jet-ci/internal/user/domain/userinfo"
	handler "jet-ci/internal/user/handler/user"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	utils.Api
	UserHandler handler.IUserHandler
}

func NewUserApi() UserApi {
	return UserApi{
		UserHandler: handler.NewUserHandler(),
	}
}

func UserRouting(r *gin.RouterGroup) {
	api := NewUserApi()
	r.POST("/user/add", api.AddUser)

}

func (api *UserApi) AddUser(ctx *gin.Context) {
	var input userinfo.InputCreateUser
	api.MakeContext(ctx, &input)

	err := api.UserHandler.CreateUser(&input)
	if err != nil {
		api.Error(err.Error())
	} else {
		api.OkMsg("添加成功!")
	}
}

// func (api *UserApi) login(ctx *gin.Context) {
// 	api.MakeContext(ctx)

// 	nums := [5]int{1, 2, 3, 4, 5}
// 	api.Page(nums, 10, 3, 5, "OK")
// 	// ctx.AbortWithStatusJSON(http.StatusOK, res)

// 	//ctx.JSON(http.StatusOK, gin.H{"data": 1111})

// 	api.Error("出错了")

// 	// api.Error("出错了2")

// 	//api.OkMsg("Ok")  gin.H

// 	// res := A1{a1: "1111111111111", b1: "2222222222222222"}

// 	// fmt.Println("----------------->res=", res)

// 	// ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"res": res})

// 	//api.Page(res, 10, 3, 5, "OK")

// 	// ctx.JSON(200, &ResponseResult{
// 	// 	Result: false,
// 	// 	Msg:    "111",
// 	// 	Code:   200,
// 	// 	Data:   msg,
// 	// })

// 	// res := &A1{
// 	// 	C1: "1111111111111",
// 	// 	D1: "66666655",
// 	// }
// 	// res.SD = S1{
// 	// 	Id:   100,
// 	// 	Name: "Jet",
// 	// }

// 	u := &userinfo.Account{}
// 	res := u.New()

// 	api.OK(res)
// }

// type A1 struct {
// 	C1 string `json:"c1"`
// 	D1 string `json:"d1"`
// 	SD S1
// }

// type S1 struct {
// 	Id   int
// 	Name string
// }
