package routers

import (
	"livingserver/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// add by yyff
	beego.Router("/api/emotion", &controllers.EmotionController{}, "get:GetAllEmotion")
	beego.Router("/api/emotion", &controllers.EmotionController{}, "post:PostEmotion")
	beego.Router("/api/emotion/self", &controllers.EmotionController{}, "get:GetEmotionByUser")
	beego.Router("/api/like", &controllers.LikeController{}, "post:PostLike")

	// Added by Wiesenyang
	beego.Router("/api/user", &controllers.UserController{}, "POST:Register")
	beego.Router("/api/user/login", &controllers.UserController{}, "POST:Login")
	beego.Router("/api/user/logout", &controllers.UserController{}, "GET:Logout")

	beego.Router("/api/comment", &controllers.CommentController{}, "POST:PostComment")
	beego.Router("/api/comment", &controllers.CommentController{}, "GET:GetAllComment")
	beego.Router("/api/message", &controllers.MessageController{}, "GET:GetAllMessage")
	beego.Router("/api/emotion/this", &controllers.EmotionController{}, "GET:GetEmotionById")

	beego.Router("/api/user", &controllers.UserController{}, "get:GetUserInfo")

}
