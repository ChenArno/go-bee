package base

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"log"
)

// 无需登录
var NOTFILTERROUTER = []string{
	"/v1/user/login",
}

type FilterStruct struct {
	ResponseModel
}

// 登录状态过滤
func (f *FilterStruct) FilterLoginStatus() beego.FilterFunc {
	filter := func(ctx *context.Context) {
		f.Ctx = ctx
		// 登录session是否存在
		ok := ctx.Input.Header("token")

		uri := ctx.Request.URL.Path
		log.Print("uri:" + uri)
		exist := IsContainInList(NOTFILTERROUTER, uri)
		if ok == "" && !exist {
			f.ResponseResult(CodeMessage("WITHOUTLOGIN"), "token 不存在")
		}
	}
	return filter
}

func IsContainInList(items []string, item string) bool {
	for _, v := range items {
		if item == v {
			return true
		}
	}
	return false
}
