package routers

import (
	"livingserver/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// Added by Wiesenyang
	beego.Router("/api/user", &controllers.UserController{}, "POST:Register")
	beego.Router("/api/user/login", &controllers.UserController{}, "POST:Login")
	beego.Router("/api/user/logout", &controllers.UserController{}, "GET:Logout")

	beego.Router("/api/comment", &controllers.CommentController{}, "POST:PostComment")
	beego.Router("/api/comment", &controllers.CommentController{}, "GET:GetAllComment")
	beego.Router("/api/message", &controllers.MessageController{}, "GET:GetAllMessage")
	beego.Router("/api/emotion", &controllers.EmotionController{}, "GET:GetEmotionById")

}
