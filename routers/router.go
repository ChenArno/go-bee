package routers

import (
	"bee-shua/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/user", &controllers.UserController{})
	//beego.Router("/cms", &controllers.CmsController{}, "get:GetFunc;post:PostFunc")
	userRouter()
}
