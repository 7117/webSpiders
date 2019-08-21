package routers

import (
	"web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{},"get:Get")
	beego.Router("/", &controllers.MainController{},"post:Post")

}
