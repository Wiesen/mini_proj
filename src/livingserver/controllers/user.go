package controllers

import (
	"mini_proj/src/livingserver/models"
	"github.com/astaxie/beego"
	"mini_proj/src/livingserver/filters"
	"regexp"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

func (c *UserController) Detail() {
	id := c.Ctx.Input.Param(":id")
	ok, user := models.FindUserById(id)
	if ok {
		c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
		c.Data["PageTitle"] = "Personal"
		c.Data["CurrentUserInfo"] = user
		c.Data["Topics"] = models.FindPostByUser(&user, 10)
		//c.Data["Replies"] = models.FindReplyByUser(&user, 7)
	}
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/detail.tpl"
}

func (c *UserController) SettingPage() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	c.Data["PageTitle"] = "Setting"
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/setting.tpl"
}

func (c *UserController) Setting() {
	flash := beego.NewFlash()
	id, username := c.Input().Get("id"), c.Input().Get("username")
	if len(id) > 0 {
		ok, _ := regexp.MatchString("^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\\d{8}$", id)
		if !ok {
			flash.Error("Please input correct phone number")
			flash.Store(&c.Controller)
			c.Redirect("/user/setting", 302)
			return
		}
	}
	if len(username) > 10 {
		flash.Error("Username is too long", username)
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}
	// todo: username判重
	// flash.Error("%d already existed", username)
	_, user := filters.IsLogin(c.Ctx)
	user.Id = id
	user.Username = username
	models.UpdateUser(&user)
	flash.Success("Setting update succeeded")
	flash.Store(&c.Controller)
	c.Redirect("/user/setting", 302)
}

func (c *UserController) UpdatePwd() {
	flash := beego.NewFlash()
	oldpwd, newpwd := c.Input().Get("oldpwd"), c.Input().Get("newpwd")
	_, user := filters.IsLogin(c.Ctx)
	if user.Password != oldpwd {
		flash.Error("Incorrect password")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}
	if len(newpwd) == 0 {
		flash.Error("Password can not be NULL")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}
	user.Password = newpwd
	models.UpdateUser(&user)
	flash.Success("Password update succeeded")
	flash.Store(&c.Controller)
	c.Redirect("/user/setting", 302)
}

//func (c *UserController) UpdateAvatar() {
//	flash := beego.NewFlash()
//	f, h, err := c.GetFile("avatar")
//	if err == http.ErrMissingFile {
//		flash.Error("Please choose your avatar")
//		flash.Store(&c.Controller)
//		c.Redirect("/user/setting", 302)
//	}
//	defer f.Close()
//	if err != nil {
//		flash.Error("Failed to upload")
//		flash.Store(&c.Controller)
//		c.Redirect("/user/setting", 302)
//		return
//	} else {
//		c.SaveToFile("avatar", "static/upload/avatar/" + h.Filename)
//		_, user := filters.IsLogin(c.Ctx)
//		user.Avatar = "/static/upload/avatar/" + h.Filename
//		models.UpdateUser(&user)
//		flash.Success("Upload succeeded")
//		flash.Store(&c.Controller)
//		c.Redirect("/user/setting", 302)
//	}
//}