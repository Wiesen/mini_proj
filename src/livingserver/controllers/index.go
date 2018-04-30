package controllers

import (
	"github.com/astaxie/beego"
	"mini_proj/src/livingserver/models"
	"github.com/sluu99/uuid"
	"mini_proj/src/livingserver/filters"
)

type IndexController struct {
	beego.Controller
}

// Get: home
func (c *IndexController) Index() {
	c.Ctx.Redirect(302, "/login")
}

// Get: login
func (c *IndexController) LoginPage() {
	IsLogin, _ := filters.IsLogin(c.Ctx)
	if IsLogin {
		c.Redirect("/RecordPage", 302)
	} else {
		beego.ReadFromRequest(&c.Controller)
		c.Data["PageTitle"] = "Login"
		c.Layout = "layout/layout.tpl"
		c.TplName = "login.tpl"
	}
}

// Post: login
func (c *IndexController) Login() {
	flash := beego.NewFlash()
	id, password := c.Input().Get("id"), c.Input().Get("password")
	if flag, user := models.LoginByID(id, password); flag {
		c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), user.Token, 30 * 24 * 60 * 60, "/", beego.AppConfig.String("cookie.domain"), false, true)
		c.Redirect("/RecordPage", 302)
	} else {
		flash.Error("Invalid phone_number or password")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
	}
}

// Get: register
func (c *IndexController) RegisterPage() {
	IsLogin, _ := filters.IsLogin(c.Ctx)
	if IsLogin {
		c.Redirect("/RecordPage", 302)
	} else {
		beego.ReadFromRequest(&c.Controller)
		c.Data["PageTitle"] = "Register"
		c.Layout = "layout/layout.tpl"
		c.TplName = "register.tpl"
	}
}

// Post: register
func (c *IndexController) Register() {
	flash := beego.NewFlash()
	id, password := c.Input().Get("id"), c.Input().Get("password")
	if len(id) == 0 || len(password) == 0 {
		flash.Error("id or password can not be NULL")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
	} else if flag, _ := models.FindUserById(id); flag {
		flash.Error("id already existed")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
	} else {
		var token = uuid.Rand().Hex()
		user := models.User{Id: id, Password: password, Avatar: 0, Token: token}
		models.AddUser(&user)
		// others are ordered as cookie's max age time, path,domain, secure and http only.
		c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), token, 30 * 24 * 60 * 60, "/", beego.AppConfig.String("cookie.domain"), false, true)
		c.Redirect("/", 302)
	}
}

// Get: logout
func (c *IndexController) Logout() {
	c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), "", -1, "/", beego.AppConfig.String("cookie.domain"), false, true)
	c.Redirect("/", 302)
}

// Get: about
func (c *IndexController) About() {
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
	c.Data["PageTitle"] = "About"
	c.Layout = "layout/layout.tpl"
	c.TplName = "about.tpl"
}