package main

import (
	_ "bee-shua/routers"
	"bee-shua/routers/base"
	"github.com/astaxie/beego"
)

func main() {
	//beego.BConfig.Listen.Graceful = true
	// 静态文件
	beego.SetStaticPath("/public", "public")
	// 路由过滤
	filter := &base.FilterStruct{}
	beego.InsertFilter("/v1/*", beego.BeforeRouter, filter.FilterLoginStatus())

	beego.Run()
}
