package controllers

import (
	"encoding/json"
	"fmt"
	"livingserver/models"
	"livingserver/redis_client"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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

// the following is added by yyff
func (c *EmotionController) GetEmotionByUser() {
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

	// 获取心情列表
	isErr, emotions := models.GetEmotionByUser(user.Id, pageNo)
	if isErr {
		rsp.RetCode = -1
		rsp.Message = fmt.Sprintf("query 'emotion' failed, user id: [%v]", user.Id)
		return
	}

	// 构造响应
	for i := 0; i < len(emotions); i++ {
		m := make(map[string]interface{})
		m["emotion_id"] = emotions[i].Id
		m["content"] = emotions[i].Content
		m["label_id"] = emotions[i].Label.Id
		m["label_name"] = emotions[i].Label.LabelName
		m["strong"] = emotions[i].Strong
		m["visiable"] = emotions[i].Visiable
		m["create_time"] = emotions[i].CreateTime
		rsp.Data = append(rsp.Data, m)
	}
	logs.Info("Get %v emotions by user[%v]", len(emotions), user.Id)
}

func (c *EmotionController) GetEmotionById() {
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
	emotion_id, _ := c.GetInt("emotion_id")
	emotion, err := models.GetEmotionById(emotion_id)
	if err != nil {
		rsp.RetCode = -2
		rsp.Message = fmt.Sprintf("Invalid emotion_id")
		return
	}

	if emotion.Poster.Id != user.Id && emotion.Visiable == 1 {
		rsp.RetCode = -2
		rsp.Message = fmt.Sprintf("Invalid: you have no Authority to access this emotion")
		return
	}

	// 构造响应
	m := make(map[string]interface{})
	m["emotion_id"] = emotion.Id
	m["content"] = emotion.Content
	m["label_id"] = emotion.Label.Id
	m["label_name"] = emotion.Label.LabelName
	m["strong"] = emotion.Strong
	m["visiable"] = emotion.Visiable
	m["create_time"] = emotion.CreateTime
	rsp.Data = append(rsp.Data, m)
	logs.Info("Get emotion[%v] by user[%v]", m, user.Id)
}

func (c *EmotionController) GetAllEmotion() {

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
		return
	}

	// 获取用户点赞列表
	isErr, likes := models.GetLikeByUser(user.Id)
	if isErr {
		rsp.RetCode = -1
		rsp.Message = fmt.Sprintf("query 'like' failed")
		return
	}
	// fmt.Println("retrive like list, length:", len(likes))

	// 建立点赞查询map
	likeMap := make(map[int]int)
	for i := 0; i < len(likes); i++ {
		eid := likes[i].Emotion.Id
		likeMap[eid] = 1
	}

	// fmt.Println("start to encapsulate response package")
	// 构造响应
	for i := 0; i < len(emotions); i++ {
		m := make(map[string]interface{})
		m["emotion_id"] = emotions[i].Id
		m["content"] = emotions[i].Content
		m["label_id"] = emotions[i].Label.Id
		if label, err := models.GetLabelById(emotions[i].Label.Id); err == nil {
			m["label_name"] = label.LabelName
		} else {
			m["label_name"] = ""
		}
		m["strong"] = emotions[i].Strong
		m["create_time"] = emotions[i].CreateTime
		m["poster"] = emotions[i].Poster.Id
		if u, err := models.GetUserById(emotions[i].Poster.Id); err == nil { // fix bug: get user info
			m["nickname"] = u.Nickname
			m["avatar"] = u.Avatar
		} else {
			m["nickname"] = ""
			m["avatar"] = ""
		}
		// m["like_cnt"] = emotions[i].LikeCnt
		// m["comment_cnt"] = emotions[i].CommentCnt

		// 获取点赞数和评论数
		lc, cc := redis_client.GetLikeCnt(emotions[i].Id), redis_client.GetCommentCnt(emotions[i].Id)
		if lc < 0 || cc < 0 {
			rsp.RetCode = -1
			rsp.Message = fmt.Sprintf("get counter from redis failed, emotion id: ", emotions[i].Id)
			rsp.Data = []interface{}{}
			return
		}
		m["like_cnt"] = lc
		m["comment_cnt"] = cc

		// 判断用户是否点过赞
		if _, ok := likeMap[emotions[i].Id]; ok {
			m["is_like"] = 1
		} else {
			m["is_like"] = 0
		}
		// fmt.Println("response pack:", i, m["emotion_id"], m["content"], m["label_id"], m["label_name"], m["strong"],
		// 	m["create_time"], m["poster"], m["nickname"], m["avatar"])
		// logs.Info("Response emotion[%v]: %v", i, m)
		rsp.Data = append(rsp.Data, m)
	}
	logs.Info("Get %v emotions by user[%v]", len(emotions), user.Id)

}

func (c *EmotionController) PostEmotion() {
	// var v models.Emotion
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

	// 获取请求
	var req PostEmotionReq
	beego.ReadFromRequest(&c.Controller)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		rsp.RetCode = -1
		rsp.Message = fmt.Sprint("parse request parameter failed, request body: ", string(c.Ctx.Input.RequestBody))
		return
	}

	// 构造心情
	v := models.Emotion{
		Content: req.Content,
		Label: &models.Label{
			Id: req.LabelID,
		},
		Strong:     req.Strong,
		CreateTime: time.Now(),
		Visiable:   req.Visiable,
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
		return
	}
	logs.Info("Add emotion successful, info: %+v ", req)
}
