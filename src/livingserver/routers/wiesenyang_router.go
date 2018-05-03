package routers

import (
	"github.com/astaxie/beego"
	"livingserver/controllers"
)

func init() {
	// Added by Wiesenyang
	beego.Router("/api/user/login", &controllers.UserController{}, "POST:Login")
	beego.Router("/api/user/logout", &controllers.UserController{}, "GET:Logout")
}
