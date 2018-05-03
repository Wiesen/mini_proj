package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["livingserver/controllers:CommentController"] = append(beego.GlobalControllerRouter["livingserver/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:CommentController"] = append(beego.GlobalControllerRouter["livingserver/controllers:CommentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:CommentController"] = append(beego.GlobalControllerRouter["livingserver/controllers:CommentController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:CommentController"] = append(beego.GlobalControllerRouter["livingserver/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:CommentController"] = append(beego.GlobalControllerRouter["livingserver/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:EmotionController"] = append(beego.GlobalControllerRouter["livingserver/controllers:EmotionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:EmotionController"] = append(beego.GlobalControllerRouter["livingserver/controllers:EmotionController"],
		beego.ControllerComments{
			Method: "GetAllEmotion",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:EmotionController"] = append(beego.GlobalControllerRouter["livingserver/controllers:EmotionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:EmotionController"] = append(beego.GlobalControllerRouter["livingserver/controllers:EmotionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:EmotionController"] = append(beego.GlobalControllerRouter["livingserver/controllers:EmotionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:EmotionController"] = append(beego.GlobalControllerRouter["livingserver/controllers:EmotionController"],
		beego.ControllerComments{
			Method: "GetEmotionByUser",
			Router: `/self`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:LabelController"] = append(beego.GlobalControllerRouter["livingserver/controllers:LabelController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:LabelController"] = append(beego.GlobalControllerRouter["livingserver/controllers:LabelController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:LabelController"] = append(beego.GlobalControllerRouter["livingserver/controllers:LabelController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:LabelController"] = append(beego.GlobalControllerRouter["livingserver/controllers:LabelController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:LabelController"] = append(beego.GlobalControllerRouter["livingserver/controllers:LabelController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:LikeController"] = append(beego.GlobalControllerRouter["livingserver/controllers:LikeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:LikeController"] = append(beego.GlobalControllerRouter["livingserver/controllers:LikeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:LikeController"] = append(beego.GlobalControllerRouter["livingserver/controllers:LikeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:LikeController"] = append(beego.GlobalControllerRouter["livingserver/controllers:LikeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:LikeController"] = append(beego.GlobalControllerRouter["livingserver/controllers:LikeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["livingserver/controllers:UserController"] = append(beego.GlobalControllerRouter["livingserver/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
