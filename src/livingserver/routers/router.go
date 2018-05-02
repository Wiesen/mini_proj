// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"livingserver/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/comment",
			beego.NSInclude(
				&controllers.CommentController{},
			),
		),

		beego.NSNamespace("/emotion",
			beego.NSInclude(
				&controllers.EmotionController{},
			),
		),

		beego.NSNamespace("/label",
			beego.NSInclude(
				&controllers.LabelController{},
			),
		),

		beego.NSNamespace("/like",
			beego.NSInclude(
				&controllers.LikeController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)

	// Added by Wiesenyang
	beego.Router("/v1/user/login", &controllers.UserController{}, "POST:Login")
	beego.Router("/v1/user/logout/?:token", &controllers.UserController{}, "POST:Logout")
}
