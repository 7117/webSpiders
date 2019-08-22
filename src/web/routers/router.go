package routers

import (
	"web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{},"get:Test")
	beego.Router("/", &controllers.MainController{},"post:Post")
	beego.Router("/test", &controllers.TestController{}, "get:Get;post:Post")
	beego.Router("/testinput", &controllers.TestInputController{}, "get:Get;post:Post")
	beego.Router("/testlogin", &controllers.TestLoginController{}, "get:Login;post:Post")
	beego.Router("/testmodel", &controllers.TestModelController{}, "get:Get;post:Post")

}
