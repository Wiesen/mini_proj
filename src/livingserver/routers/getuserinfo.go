package routers

import (
	"livingserver/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/user", &controllers.UserController{}, "get:GetUserInfo")
}
