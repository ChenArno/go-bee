package routers

import (
	"bee-shua/controllers"
	"github.com/astaxie/beego"
)

func userRouter() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSRouter("/login", &controllers.UserController{}, "post:Post"),
			beego.NSRouter("/getAllUserList", &controllers.UserController{}, "*:GetAll")))

	beego.AddNamespace(ns)
}
