package controllers

import "github.com/astaxie/beego"

type CmsController struct {
	beego.Controller
}

func (c *CmsController) Get() {
	c.Data["json"] = map[string]interface{}{"result": "ss", "msg": "获取成功", "code": "200"} // 设置返回值
	c.ServeJSON()
}

func (c *CmsController) Post() {
	c.Data["json"] = map[string]interface{}{"result": "ss", "msg": "获取成功", "code": "200"} // 设置返回值
	c.ServeJSON()
}
