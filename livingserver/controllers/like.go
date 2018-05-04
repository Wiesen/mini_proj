package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"fmt"
	"time"
	"github.com/Wiesen/mini_proj/livingserver/models"

	"github.com/astaxie/beego"
)

// LikeController operations for Like
type LikeController struct {
	beego.Controller
}

// URLMapping ...
func (c *LikeController) URLMapping() {
	// c.Mapping("Post", c.Post)
	// c.Mapping("GetOne", c.GetOne)
	// c.Mapping("GetAll", c.GetAll)
	// c.Mapping("Put", c.Put)
	// c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Like
// @Param	body		body 	models.Like	true		"body for Like content"
// @Success 201 {int} models.Like
// @Failure 403 body is empty
// @router / [post]
func (c *LikeController) Post() {
	var v models.Like
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddLike(&v); err == nil {
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
// @Description get Like by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Like
// @Failure 403 :id is empty
// @router /:id [get]
func (c *LikeController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetLikeById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Like
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Like
// @Failure 403
// @router / [get]
func (c *LikeController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllLike(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Like
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Like	true		"body for Like content"
// @Success 200 {object} models.Like
// @Failure 403 :id is not int
// @router /:id [put]
func (c *LikeController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Like{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateLikeById(&v); err == nil {
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
// @Description delete the Like
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *LikeController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteLike(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}



// the following is added by yyff
func (c *LikeController) PostLike() {
	rsp := CommonRsp{RetCode: 0}

	defer func() {
		c.Data["json"] = rsp
		c.ServeJSON()
	}()

	token := c.GetString("token")
	hasRows, user := models.GetUserByToken(token)
	if !hasRows {
		rsp.RetCode = -2
		rsp.Message = fmt.Sprintf("Invalid token")
		return
	}

	// inputMap := make(map[string]interface{})
	var req PostLikeReq
	beego.ReadFromRequest(&c.Controller)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		rsp.RetCode = -1
		rsp.Message = fmt.Sprint("parse request parameter failed, request body: ", string(c.Ctx.Input.RequestBody))
		return
	}

	v := models.Like {
		EmotionId : &models.Emotion{
			Id : req.EmotionID,
		},
		Poster: &models.User{
			Id: user.Id,
		},
		CreateTime : time.Now(),
	}
	
	_, err = models.AddLike(&v)
	if err != nil {
		rsp.RetCode = -1
		rsp.Message = err.Error() 
		return
	}
	
}