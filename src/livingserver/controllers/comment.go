package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"livingserver/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// CommentController operations for Comment
type CommentController struct {
	beego.Controller
}

func (c *CommentController) PostComment() {
	rsp := CommonRsp{RetCode: 0}

	defer func() {
		c.Data["json"] = rsp
		c.ServeJSON()
	}()

	// 获取url参数
	token := c.GetString("token")
	hasRows, user := models.GetUserByToken(token)
	if !hasRows {
		rsp.RetCode = -2
		rsp.Message = fmt.Sprintf("Invalid token")
		return
	}

	// inputMap := make(map[string]interface{})
	var req PostCommentReq
	beego.ReadFromRequest(&c.Controller)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		rsp.RetCode = -1
		rsp.Message = fmt.Sprint("parse request parameter failed, request body: ", string(c.Ctx.Input.RequestBody))
		return
	}
	v := models.Comment{
		Content:    req.Comment,
		Emotion:    &models.Emotion{Id: req.EmotionID},
		Poster:     &models.User{Id: user.Id},
		CreateTime: time.Now(),
		Rspto:      req.RspTo,
	}

	if _, err := models.AddComment(&v); err != nil {
		rsp.RetCode = -1
		rsp.Message = err.Error()
		return
	}

	// added by wiesenyang
	m := models.Message{
		CreateTime: v.CreateTime,
		Emotion:    v.Emotion,
		Content:    v.Content,
		Poster:     v.Poster,
	}
	emt, _ := models.GetEmotionById(v.Emotion.Id)
	if v.Rspto == 0 {
		m.TypeId = 2
		m.Owner = emt.Poster
	} else {
		m.TypeId = 3
		m.Owner = &models.User{Id: v.Rspto}
	}
	if _, err := models.AddMessage(&m); err != nil {
		rsp.RetCode = -1
		rsp.Message = err.Error()
		return
	}
	logs.Info("add Comment successful, info:", m.CreateTime, m.Emotion, m.Content, m.Poster, m.Owner, m.TypeId)
}

// @router / [get]
func (c *CommentController) GetAllComment() {

	rsp := CommonRsp{RetCode: 0}
	defer func() {
		c.Data["json"] = rsp
		c.ServeJSON()
	}()

	// 获取url参数
	token := c.GetString("token")
	hasRows, _ := models.GetUserByToken(token)
	if !hasRows {
		rsp.RetCode = -2
		rsp.Message = fmt.Sprintf("Invalid token")
		return
	}
	pageNo, err := c.GetInt("pageno")
	if err != nil || pageNo < 0 {
		rsp.RetCode = -1
		rsp.Message = fmt.Sprintf("Invalid pageno")
		return
	}

	// 获取comment列表
	var isErr bool
	var comments []*models.Comment
	if v, err := c.GetInt("emotion_id"); err == nil {
		isErr, comments = models.GetCommentByEmotion(v, pageNo)
		if isErr {
			rsp.RetCode = -1
			rsp.Message = fmt.Sprintf("query 'comment' failed")
			return
		}
	} else {
		rsp.RetCode = -1
		rsp.Message = fmt.Sprintf("Invalid request json: no emotion id")
		return
	}

	// 构造响应
	for i := 0; i < len(comments); i++ {
		m := make(map[string]interface{})
		m["comment_id"] = comments[i].Id
		m["comment"] = comments[i].Content
		m["poster"] = comments[i].Poster.Id
		u, _ := models.GetUserById(comments[i].Poster.Id)
		m["poster_nickname"] = u.Nickname
		m["post_avatar"] = u.Avatar
		m["create_time"] = comments[i].CreateTime.Format("2006-01-02 15:04:05")
		m["rspto"] = comments[i].Rspto
		if comments[i].Rspto != 0 {
			if user, err := models.GetUserById(comments[i].Rspto); err == nil {
				m["rspto_nickname"] = user.Nickname
			}
		}

		rsp.Data = append(rsp.Data, m)
	}
}
