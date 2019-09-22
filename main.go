package main

import (
	"nvwa-console-api/controller"
	model "nvwa-console-api/model"
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/logger"
)

/**
 * 初始化 iris app
 * @method NewApp
 * @return  {[type]}  api      *iris.Application  [iris app]
 */
func newApp() (api *iris.Application) {
	api = iris.New()
	api.Use(logger.New())

	api.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(model.APIResponse(false, nil, "404 Not Found"))
	})
	api.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.WriteString("Oups something went wrong, try again")
	})

	v1 := api.Party("/v1").AllowMethods(iris.MethodOptions)
	{
		//v1.Use(middleware.NewYaag()) // <- IMPORTANT, register the middleware.
		v1.Post("/admin/login", controller.UserLogin)
		v1.Get("/admin/login", controller.UserLogin)
		// v1.PartyFunc("/admin", func(admin router.Party) {
		// 	admin.Use(middleware.JwtHandler().Serve, middleware.AuthToken)
		// 	admin.Get("/logout", controllers.UserLogout)
		// })
	}

	return
}

func main() {
	app := newApp()

	addr := ":9001" //config.Conf.Get("app.addr").(string)
	app.Run(iris.Addr(addr))
}
