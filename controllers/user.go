package controllers

import (
	"bee-shua/routers/base"
	"encoding/json"
	"fmt"
)

type UserController struct {
	base.ResponseModel
}
type MovieInfo struct {
	Id   string `json:"id"`   // 注意首字母大写
	Name string `json:"name"` // 并且写明与json字段的映射关系，否则Unmarshal函数不好用
}

func (u *UserController) Get() {

	var MovieinfoObj MovieInfo
	Params := u.Ctx.Input.Query("id") //这是request的body 的json二进制数据
	fmt.Println(Params)
	json.Unmarshal([]byte(`{
		"name":"ss",
		"id":"ss"
	}`), &MovieinfoObj) //解析二进制json，把结果放进MovieInfo_obj中
	fmt.Println(MovieinfoObj)

	u.Data["json"] = map[string]interface{}{"result": MovieinfoObj, "msg": "获取成功", "code": "200"} // 设置返回值
	u.ServeJSON()                                                                                 // 返回json数据
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	modelDemo.User	true		"body for user content"
// @Success 200 {int} modelDemo.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	json.Unmarshal(u.Ctx.Input.RequestBody, "user")
	u.ResponseResult(base.CodeMessage("OK"), "uid")
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} modelDemo.User
// @router / [get]
func (u *UserController) GetAll() {

	u.ResponseResult(base.CodeMessage("OK"), "users")
}
