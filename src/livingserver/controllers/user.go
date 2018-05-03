package controllers

import (
	"encoding/json"
	_ "errors"
	"livingserver/models"
	_ "strconv"
	_ "strings"
	"livingserver/filters"
	"github.com/astaxie/beego"
	"github.com/sluu99/uuid"
)

// UserController operations for User
type UserController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	//c.Mapping("GetOne", c.GetOne)
	//c.Mapping("GetAll", c.GetAll)
	//c.Mapping("Put", c.Put)
	//c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	//var v models.User
	input_table, output_table := make(map[string]interface{}), make(map[string]interface{})
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &input_table); err == nil {
		phone_number, password, nickname := input_table["phone_number"], input_table["password"], input_table["nickname"]
		if phone_number == nil || password == nil || nickname == nil {
			return
		}
		user := models.User{PhoneNumber:phone_number.(string), Password:password.(string), Nickname: nickname.(string)}
		if _, err := models.AddUser(&user); err == nil {
			user.Token  = uuid.Rand().Hex()
			models.UpdateUserById(&user)
			output_table["ret_code"] = 0
			output_table["data"] = map[string]string{"token": user.Token}
			//c.Ctx.Output.SetStatus(201)
		} else {
			output_table["ret_code"] = -1
			output_table["message"] = "database create user failed"
		}
	} else {
		output_table["ret_code"] = -1
		output_table["message"] = "Invalid json format"
	}
	c.Data["json"] = output_table
	c.ServeJSON()
}

//// Put ...
//// @Title Put
//// @Description update the User
//// @Param	id		path 	string	true		"The id you want to update"
//// @Param	body		body 	models.User	true		"body for User content"
//// @Success 200 {object} models.User
//// @Failure 403 :id is not int
//// @router /:id [put]
//func (c *UserController) Put() {
//	idStr := c.Ctx.Input.Param(":id")
//	id, _ := strconv.Atoi(idStr)
//	v := models.User{Id: id}
//	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
//		if err := models.UpdateUserById(&v); err == nil {
//			c.Data["json"] = "OK"
//		} else {
//			c.Data["json"] = err.Error()
//		}
//	} else {
//		c.Data["json"] = err.Error()
//	}
//	c.ServeJSON()
//}

//// Delete ...
//// @Title Delete
//// @Description delete the User
//// @Param	id		path 	string	true		"The id you want to delete"
//// @Success 200 {string} delete success!
//// @Failure 403 id is empty
//// @router /:id [delete]
//func (c *UserController) Delete() {
//	idStr := c.Ctx.Input.Param(":id")
//	id, _ := strconv.Atoi(idStr)
//	if err := models.DeleteUser(id); err == nil {
//		c.Data["json"] = "OK"
//	} else {
//		c.Data["json"] = err.Error()
//	}
//	c.ServeJSON()
//}


// Post: login
func (c *UserController) Login() {
	input_table, output_table := make(map[string]interface{}), make(map[string]interface{})
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &input_table); err == nil {
		phone_number, password := input_table["phone_number"], input_table["password"]
		if phone_number == nil || password == nil {
			return
		}
		if flag, user := models.Login(phone_number.(string), password.(string)); flag {
			user.Token = uuid.Rand().Hex()
			models.UpdateUserById(&user)
			output_table["ret_code"] = 0
			output_table["data"] = map[string]string{"token": user.Token}
		} else {
			flash.Error("Invalid phone_number or password")
			flash.Store(&c.Controller)
			output_table["ret_code"], output_table["message"] = -1, "Invalid phone_number or password"
		}
	} else {
		flash.Error("Invalid json format")
		flash.Store(&c.Controller)
		output_table["ret_code"], output_table["message"] = -1, "Invalid json data"
	}
	c.Data["json"] = output_table
	c.ServeJSON()
}

// Post: logout
func (c *UserController) Logout() {
	table := make(map[string]interface{})
	beego.ReadFromRequest(&c.Controller)
	if err, user := filters.IsLogin(c.Ctx); err == true {
		user.Token = ""
		models.UpdateUserById(&user)
		table["ret_code"] = 0
	} else {
		table["ret_code"] = -2
		table["message"] = "Invalid token"
	}
	c.Data["json"] = table
	c.ServeJSON()
}