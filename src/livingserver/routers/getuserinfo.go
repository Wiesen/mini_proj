package routers

import (
	"github.com/astaxie/beego"
	"livingserver/controllers"
)
func init() {
	beego.Router("/api/user",&controllers.UserController{},"get:GetUserInfo")
}