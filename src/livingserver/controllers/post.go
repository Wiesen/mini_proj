package controllers

import (
	"github.com/astaxie/beego"
	"mini_proj/src/livingserver/models"
	"strconv"
	"mini_proj/src/livingserver/filters"
)

type PostController struct {
	beego.Controller
}

func (c *PostController) RecordPage() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
	c.Data["PageTitle"] = "Record emotion is this moment"
	c.Data["Sections"] = models.FindAllSection()
	c.Layout = "layout/layout.tpl"
	c.TplName = "Post/create.tpl"
}

func (c *PostController) Record() {
	flash := beego.NewFlash()
	content, section := c.Input().Get("content"), c.Input().Get("section")
	if len(content) > 120 {
		flash.Error("Content can not exceed more than 120 characters")
		flash.Store(&c.Controller)
		c.Redirect("/Post/create", 302)
	} else {
		s, _ := strconv.Atoi(section)
		if s == 0 {
			flash.Error("Please choose your emotion section")
			flash.Store(&c.Controller)
			c.Redirect("/Post/create", 302)
		} else {
			section := models.Section{Id: s}
			_, user := filters.IsLogin(c.Ctx)
			Post := models.Post{ Content: content, Section: &section, User: &user}
			id := models.AddPost(&Post)
			c.Redirect("/Post/" + strconv.FormatInt(id, 10), 302)
		}
	}
}

func (c *PostController) Detail() {
	id := c.Ctx.Input.Param(":id")
	tid, _ := strconv.Atoi(id)
	if tid > 0 {
		c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
		Post := models.FindPostById(tid)
		models.IncrViewCount(&Post)
		c.Data["Post"] = Post
		c.Data["Replies"] = models.FindReplyByPost(&Post)
		c.Layout = "layout/layout.tpl"
		c.TplName = "Post/detail.tpl"
	} else {
		c.Ctx.WriteString("Post doesn't exist")
	}
}

func (c *PostController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		Post := models.FindPostById(id)
		models.DeletePost(&Post)
		models.DeleteReplyByPost(&Post)
		c.Redirect("/", 302)
	} else {
		c.Ctx.WriteString("Post doesn't exist")
	}
}