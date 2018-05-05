package controllers


import (
	"github.com/astaxie/beego"
	"livingserver/models"
	"fmt"
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
		m["create_time"] = messages[i].CreateTime
		m["type"] = messages[i].TypeId
		m["emotion_id"] = messages[i].Emotion.Id
		m["comment"] = messages[i].Content
		m["poster"] = messages[i].Poster.Id
		u, _ := models.GetUserById(messages[i].Poster.Id)
		m["nickname"] = u.Nickname
		m["avatar"] = u.Avatar
		rsp.Data = append(rsp.Data, m)
	}
}
