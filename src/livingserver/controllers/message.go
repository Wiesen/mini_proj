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
	//var isErr bool
	//var messages []*models.Message
	//if v, err := c.GetInt("emotion_id"); err == nil {
	//	isErr, comments = models.GetCommentByEmotion(v, pageNo)
	//	if isErr {
	//		rsp.RetCode = -1
	//		rsp.Message = fmt.Sprintf("query 'comment' failed")
	//		return
	//	}
	//} else {
	//	rsp.RetCode = -1
	//	rsp.Message = fmt.Sprintf("Invalid request json: no emotion id")
	//	return
	//}

	// 构造响应
	//for i := 0; i < len(comments); i++ {
	//	m := make(map[string]interface{})
	//	m["emotion_id"] = comments[i].EmotionId.Id
	//	m["comment"] = comments[i].Content
	//	m["poster"] = comments[i].Poster.Id
	//	m["poster_nickname"] = comments[i].Poster.Nickname
	//	m["create_time"] = comments[i].CreateTime
	//	m["rspto"] = comments[i].Rspto
	//	if comments[i].Rspto != 0 {
	//		if user, err := models.GetUserById(comments[i].Rspto); err == nil {
	//			m["rspto_nickname"] = user.Nickname
	//		}
	//	}
	//	rsp.Data = append(rsp.Data, m)
	//}
}
