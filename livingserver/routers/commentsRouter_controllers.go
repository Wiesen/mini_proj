package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:CommentController"],
		beego.ControllerComments{
			Method: "GetAllComment",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:EmotionController"] = append(beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:EmotionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:EmotionController"] = append(beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:EmotionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:EmotionController"] = append(beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:EmotionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:EmotionController"] = append(beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:EmotionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LabelController"] = append(beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LabelController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LabelController"] = append(beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LabelController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LabelController"] = append(beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LabelController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LabelController"] = append(beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LabelController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LabelController"] = append(beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LabelController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LikeController"] = append(beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LikeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LikeController"] = append(beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LikeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LikeController"] = append(beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LikeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LikeController"] = append(beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LikeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LikeController"] = append(beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:LikeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Wiesen/mini_proj/livingserver/controllers:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
