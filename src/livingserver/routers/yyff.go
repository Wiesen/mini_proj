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

}
