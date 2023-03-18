package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"mainproject/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/:id(.*)", &controllers.UsersController{})
}
