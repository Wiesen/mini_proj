package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Wiesen/mini_proj/livingserver/models"
	"github.com/astaxie/beego"
)

// CommentController operations for Comment
type CommentController struct {
	beego.Controller
}

func (c *CommentController) PostComment() {
	rsp := CommonRsp{RetCode: 0}
	for {
		// 获取url参数
		token := c.GetString("token")
		hasRows, user := models.GetUserByToken(token)
		if !hasRows {
			rsp.RetCode = -2
			rsp.Message = fmt.Sprintf("Invalid token")
			break
		}

		inputMap := make(map[string]interface{})
		beego.ReadFromRequest(&c.Controller)
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &inputMap); err != nil {
			rsp.RetCode = -1
			rsp.Message = fmt.Sprint("parse request parameter failed, request body: ", string(c.Ctx.Input.RequestBody))
			break
		}
		v := models.Comment{
			Content:    inputMap["comment"].(string),
			EmotionId:  &models.Emotion{Id: inputMap["emotion_id"].(int)},
			Poster:     &models.User{Id: user.Id},
			CreateTime: time.Now(),
			Rspto:      inputMap["rspto"].(int),
		}

		if _, err := models.AddComment(&v); err != nil {
			rsp.RetCode = -1
			rsp.Message = err.Error() //fmt.Sprint("add emotion failed")
			break
		}
		//SUCCESS
		break
	}

	c.Data["json"] = rsp
	c.ServeJSON()
}

// @router / [get]
func (c *CommentController) GetAllComment() {

	rsp := CommonRsp{RetCode: 0}
	for {

		// 获取url参数
		token := c.GetString("token")
		hasRows, _ := models.GetUserByToken(token)
		if !hasRows {
			rsp.RetCode = -2
			rsp.Message = fmt.Sprintf("Invalid token")
			break
		}
		pageNo, err := c.GetInt("pageno")
		if err != nil || pageNo < 0 {
			rsp.RetCode = -1
			rsp.Message = fmt.Sprintf("Invalid pageno")
			break
		}

		// 获取comment列表
		var isErr bool
		var comments []*models.Comment
		if v, err := c.GetInt("emotion_id"); err == nil {
			isErr, comments = models.GetCommentByEmotion(v, pageNo)
			if isErr {
				rsp.RetCode = -1
				rsp.Message = fmt.Sprintf("query 'comment' failed")
				break
			}
		} else {
			rsp.RetCode = -1
			rsp.Message = fmt.Sprintf("Invalid request json: no emotion id")
			break
		}

		// 构造响应
		for i := 0; i < len(comments); i++ {
			m := make(map[string]interface{})
			m["emotion_id"] = comments[i].Id
			m["comment"] = comments[i].Content
			m["poster"] = comments[i].Poster.Id
			m["poster_nickname"] = comments[i].Poster.Nickname
			m["create_time"] = comments[i].CreateTime
			m["rspto"] = comments[i].Rspto
			if comments[i].Rspto != 0 {
				if user, err := models.GetUserById(comments[i].Rspto); err == nil {
					m["rspto_nickname"] = user.Nickname
				}
			}
			rsp.Data = append(rsp.Data, m)
		}
		// SUCCESS
		break
	}

	c.Data["json"] = rsp
	c.ServeJSON()
}
