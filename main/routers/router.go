package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"main/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
