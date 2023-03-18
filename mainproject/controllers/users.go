package controllers

import beego "github.com/beego/beego/v2/server/web"

type UsersController struct {
	beego.Controller
}

func (c *UsersController) Get() {
	c.Data["ID"] = c.Ctx.Input.Param(":id")
	c.TplName = "users/index.tpl"
}
