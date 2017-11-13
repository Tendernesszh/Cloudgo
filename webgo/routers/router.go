package routers

import (
	"webgo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
		beego.Router("/user/:name", &controllers.UserController{}) //when visit /user/:name shows hello "name"
}
