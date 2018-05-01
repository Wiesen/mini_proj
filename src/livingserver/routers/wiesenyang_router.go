package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {
	// added by Wiesenyang
	beego.GlobalControllerRouter["livingserver/controllers:UserController"] = append(beego.GlobalControllerRouter["livingserver/controllers:UserController"],
		beego.ControllerComments{
			Method: "LoginPage",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})
}
