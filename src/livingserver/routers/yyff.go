package routers

import (
	"github.com/astaxie/beego"
	"livingserver/controllers"
)

func init() {
	// add by yyff
	beego.Router("/api/emotion", &controllers.EmotionController{}, "get:GetAllEmotion")
	beego.Router("/api/emotion", &controllers.EmotionController{}, "post:PostEmotion")
	beego.Router("/api/emotion/self", &controllers.EmotionController{}, "get:GetEmotionByUser")
}



