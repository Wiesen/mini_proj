package routers

import (
	"github.com/Wiesen/mini_proj/livingserver/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/user", &controllers.UserController{}, "get:GetUserInfo")
}
