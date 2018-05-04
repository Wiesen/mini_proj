package controllers

import (
	"encoding/json"
	// "errors"
	"github.com/Wiesen/mini_proj/livingserver/models"
	// "livingserver/filters"
	"strconv"
	// "strings"
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

// EmotionController operations for Emotion
type EmotionController struct {
	beego.Controller
}

// URLMapping ...
// func (c *EmotionController) URLMapping() {
// 	c.Mapping("Post", c.Post)
// 	c.Mapping("GetOne", c.GetOne)
// 	c.Mapping("GetAll", c.GetAll)
// 	c.Mapping("GetAll", c.GetAllEmotion)
// 	c.Mapping("Put", c.Put)
// 	c.Mapping("Delete", c.Delete)
// }

// @router / [post]
func (c *EmotionController) Post() {
	var v models.Emotion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddEmotion(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Emotion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Emotion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *EmotionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetEmotionById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Emotion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Emotion
// @Failure 403
// @router / [get]
// func (c *EmotionController) GetAll() {
// 	var fields []string
// 	var sortby []string
// 	var order []string
// 	var query = make(map[string]string)
// 	var limit int64 = 10
// 	var offset int64

// 	// fields: col1,col2,entity.col3
// 	if v := c.GetString("fields"); v != "" {
// 		fields = strings.Split(v, ",")
// 	}
// 	// limit: 10 (default is 10)
// 	if v, err := c.GetInt64("limit"); err == nil {
// 		limit = v
// 	}
// 	// offset: 0 (default is 0)
// 	if v, err := c.GetInt64("offset"); err == nil {
// 		offset = v
// 	}
// 	// sortby: col1,col2
// 	if v := c.GetString("sortby"); v != "" {
// 		sortby = strings.Split(v, ",")
// 	}
// 	// order: desc,asc
// 	if v := c.GetString("order"); v != "" {
// 		order = strings.Split(v, ",")
// 	}
// 	// query: k:v,k:v
// 	if v := c.GetString("query"); v != "" {
// 		for _, cond := range strings.Split(v, ",") {
// 			kv := strings.SplitN(cond, ":", 2)
// 			if len(kv) != 2 {
// 				c.Data["json"] = errors.New("Error: invalid query key/value pair")
// 				c.ServeJSON()
// 				return
// 			}
// 			k, v := kv[0], kv[1]
// 			query[k] = v
// 		}
// 	}

// 	l, err := models.GetAllEmotion(query, fields, sortby, order, offset, limit)
// 	if err != nil {
// 		c.Data["json"] = err.Error()
// 	} else {
// 		c.Data["json"] = l
// 	}
// 	c.ServeJSON()
// }

// Put ...
// @Title Put
// @Description update the Emotion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Emotion	true		"body for Emotion content"
// @Success 200 {object} models.Emotion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *EmotionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Emotion{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateEmotionById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Emotion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *EmotionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteEmotion(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// ----------------------------------------------------------
// the following is added by yyff

// @router /self [get]
func (c *EmotionController) GetEmotionByUser() {
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
		pageNo, err := c.GetInt("pageno")
		if err != nil {
			rsp.RetCode = -1
			rsp.Message = fmt.Sprintf("Invalid pageno")
			break
		}

		// ---mock---
		// pageNo := 0
		// user := models.User{Id : 1}

		// 获取心情列表
		isErr, emotions := models.GetEmotionByUser(user.Id, pageNo)
		if isErr {
			rsp.RetCode = -1
			rsp.Message = fmt.Sprintf("query 'emotion' failed, user id: [%v]", user.Id)
			break
		}

		// 构造响应
		for i := 0; i < len(emotions); i++ {
			m := make(map[string]interface{})
			m["emotion_id"] = emotions[i].Id
			m["content"] = emotions[i].Content
			m["label_id"] = emotions[i].LabelId.Id
			m["label_name"] = emotions[i].LabelId.LabelName
			m["strong"] = emotions[i].Strong
			m["create_time"] = emotions[i].CreateTime
			rsp.Data = append(rsp.Data, m)
		}

		// SUCCESS
		break
	}

	c.Data["json"] = rsp
	c.ServeJSON()
}

// @router / [get]
func (c *EmotionController) GetAllEmotion() {

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
		pageNo, err := c.GetInt("pageno")
		if err != nil || pageNo < 0 {
			rsp.RetCode = -1
			rsp.Message = fmt.Sprintf("Invalid pageno")
			break
		}

		// ---mock---
		// pageNo := 0
		// user := models.User{Id : 1}

		// 获取心情列表
		var isErr bool
		var emotions []*models.Emotion
		if v, err := c.GetInt("label_id"); err == nil {
			isErr, emotions = models.GetEmotionByLabel(v, pageNo)
		} else {
			isErr, emotions = models.GetAllEmotion(pageNo)
		}
		if isErr {
			rsp.RetCode = -1
			rsp.Message = fmt.Sprintf("query 'emotion' failed")
			break
		}

		// 获取用户点赞列表
		isErr, likes := models.GetLikeByUser(user.Id)
		if isErr {
			rsp.RetCode = -1
			rsp.Message = fmt.Sprintf("query 'like' failed")
			break
		}

		// 建立点赞查询map
		likeMap := make(map[int]int)
		for i := 0; i < len(likes); i++ {
			eid := likes[i].EmotionId.Id
			likeMap[eid] = 1
		}

		// 构造响应
		for i := 0; i < len(emotions); i++ {
			m := make(map[string]interface{})
			m["emotion_id"] = emotions[i].Id
			m["content"] = emotions[i].Content
			m["label_id"] = emotions[i].LabelId.Id
			m["label_name"] = emotions[i].LabelId.LabelName
			m["strong"] = emotions[i].Strong
			m["create_time"] = emotions[i].CreateTime
			m["poster"] = emotions[i].Poster.Id
			m["nickname"] = emotions[i].Poster.Nickname
			m["avatar"] = emotions[i].Poster.Avatar
			m["create_time"] = emotions[i].CreateTime
			m["like_cnt"] = emotions[i].LikeCnt
			m["comment_cnt"] = emotions[i].CommentCnt

			// 判断用户是否点过赞
			if _, ok := likeMap[emotions[i].Id]; ok {
				m["is_like"] = 1
			} else {
				m["is_like"] = 0
			}

			rsp.Data = append(rsp.Data, m)
		}
		// SUCCESS
		break
	}

	c.Data["json"] = rsp
	c.ServeJSON()
}

func (c *EmotionController) PostEmotion() {
	// var v models.Emotion
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

		err := json.Unmarshal(c.Ctx.Input.RequestBody, &inputMap)
		if err != nil {
			rsp.RetCode = -1
			rsp.Message = fmt.Sprint("parse request parameter failed, request body: ", string(c.Ctx.Input.RequestBody))
			break
		}
		// u := models.User { Id : inputMap["poster"].(int) }
		v := models.Emotion{
			Content: inputMap["content"].(string),
			LabelId: &models.Label{
				Id: int(inputMap["label_id"].(float64)),
			},
			Strong:     int8(inputMap["strong"].(float64)),
			CreateTime: time.Now(),
			Visiable:   int8(inputMap["visiable"].(float64)),
			Poster: &models.User{
				Id: user.Id,
			},
			CommentCnt: 0,
			LikeCnt:    0,
		}

		_, err = models.AddEmotion(&v)
		if err != nil {
			rsp.RetCode = -1
			rsp.Message = err.Error()
			break
		}
		//SUCCESS
		break

	}

	c.Data["json"] = rsp
	c.ServeJSON()
}
