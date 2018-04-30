package routers

import (
	"github.com/astaxie/beego"
	"mini_proj/src/livingserver/controllers"
	"mini_proj/src/livingserver/filters"
)

func init() {
	// home page
	beego.Router("/", &controllers.IndexController{}, "GET:Index")
	beego.Router("/login", &controllers.IndexController{}, "GET:LoginPage")
	beego.Router("/login", &controllers.IndexController{}, "POST:Login")
	beego.Router("/register", &controllers.IndexController{}, "GET:RegisterPage")
	beego.Router("/register", &controllers.IndexController{}, "POST:Register")
	beego.Router("/logout", &controllers.IndexController{}, "GET:Logout")
	beego.Router("/about", &controllers.IndexController{}, "GET:About")

	// user page
	beego.Router("/user/:username", &controllers.UserController{}, "GET:Detail")
	beego.Router("/user/setting", &controllers.UserController{}, "GET:SettingPage")
	beego.InsertFilter("/user/setting", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/user/setting", &controllers.UserController{}, "POST:Setting")
	beego.InsertFilter("/user/updatepwd", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/user/updatepwd", &controllers.UserController{}, "POST:UpdatePwd")

	// record page
	beego.InsertFilter("/post/create", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/post/create", &controllers.PostController{}, "GET:RecordPage")
	beego.Router("/post/create", &controllers.PostController{}, "POST:Record")
	//beego.Router("/post/:id([0-9]+)", &controllers.PostController{}, "GET:Detail")
	//beego.InsertFilter("/post/delete/:id([0-9]+)", beego.BeforeRouter, filters.FilterUser)
	//beego.Router("/post/delete/:id([0-9]+)", &controllers.PostController{}, "GET:Delete")

	//// square page
	//beego.InsertFilter("/reply/save", beego.BeforeRouter, filters.FilterUser)
	//beego.Router("/reply/save", &controllers.ReplyController{}, "POST:Save")
	//beego.InsertFilter("/reply/up", beego.BeforeRouter, filters.FilterUser)
	//beego.Router("/reply/up", &controllers.ReplyController{}, "GET:Up")
	//beego.InsertFilter("/reply/delete/:id([0-9]+)", beego.BeforeRouter, filters.FilterUser)
	//beego.Router("/reply/delete/:id([0-9]+)", &controllers.ReplyController{}, "GET:Delete")
}
