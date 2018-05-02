package routers

import (
	"github.com/astaxie/beego"
	"livingserver/controllers"
)

func init() {
	// Added by Wiesenyang
	beego.Router("/v1/user/login", &controllers.UserController{}, "POST:Login")
	beego.Router("/v1/user/logout/?:token", &controllers.UserController{}, "POST:Logout")
}
