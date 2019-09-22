package controller

import (
	"github.com/kataras/iris"
	model "nvwa-console-api/model"
)

func UserLogin(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(model.APIResponse(false, nil, "用户名格式错误"))
}