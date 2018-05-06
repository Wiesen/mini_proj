package controllers

import (
	"fmt"
	"livingserver/models"

	"github.com/astaxie/beego"
)

// LabelController operations for Label
type MessageController struct {
	beego.Controller
}

// @router / [get]
func (c *MessageController) GetAllMessage() {

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
	pageNo, err := c.GetInt("pageno")
	if err != nil || pageNo < 0 {
		rsp.RetCode = -1
		rsp.Message = fmt.Sprintf("Invalid pageno")
		return
	}

	// 获取message列表
	var isErr bool
	var messages []*models.Message
	isErr, messages = models.GetMessageByUser(user.Id, pageNo)
	if isErr {
		rsp.RetCode = -1
		rsp.Message = fmt.Sprintf("query 'comment' failed")
		return
	}

	// 构造响应
	for i := 0; i < len(messages); i++ {
		m := make(map[string]interface{})
		m["create_time"] = messages[i].CreateTime.Format("2006-01-02 15:04:05")
		m["type"] = messages[i].TypeId
		m["emotion_id"] = messages[i].Emotion.Id
		m["comment"] = messages[i].Content
		m["poster"] = messages[i].Poster.Id
		if u, err := models.GetUserById(messages[i].Poster.Id); err == nil {
			m["nickname"] = u.Nickname
			m["avatar"] = u.Avatar
		} else {
			m["nickname"] = ""
			m["avatar"] = ""
		}
		rsp.Data = append(rsp.Data, m)
	}
}
