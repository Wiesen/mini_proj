package controllers

import (
	"encoding/json"
	"fmt"

	"livingserver/models"
	"github.com/astaxie/beego"
	"github.com/sluu99/uuid"
)

// UserController operations for User
type UserController struct {
	beego.Controller
}

// @Title Register
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Register() {
	input_table, output_table := make(map[string]interface{}), make(map[string]interface{})
	beego.ReadFromRequest(&c.Controller)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &input_table); err == nil {
		phone_number, password, nickname, qq_number := input_table["phone_number"], input_table["password"], input_table["nickname"], input_table["qq_number"]
		if phone_number == nil || password == nil || nickname == nil || qq_number == nil {
			return
		}
		user := models.User{PhoneNumber: phone_number.(string), Password: password.(string), Nickname: nickname.(string), QqNumber: qq_number.(string)}
		fmt.Println("Create user:", user.PhoneNumber, user.Password, user.Nickname, user.QqNumber)
		if id, err := models.AddUser(&user); err == nil {
			user.Token = uuid.Rand().Hex()
			models.UpdateUserById(&user)
			output_table["ret_code"] = 0
			output_table["data"] = map[string]string{"token": user.Token}
			//c.Ctx.Output.SetStatus(201)
		} else {
			fmt.Println("Add user failed:", id, err)
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
	//flash := beego.NewFlash()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &input_table); err == nil {
		phone_number, password := input_table["phone_number"], input_table["password"]
		fmt.Println("login:", phone_number, password)
		if phone_number == nil || password == nil {
			return
		}
		if flag, user := models.Login(phone_number.(string), password.(string)); flag {
			user.Token = uuid.Rand().Hex()
			models.UpdateUserById(&user)
			output_table["ret_code"] = 0
			output_table["data"] = map[string]string{"token": user.Token}
		} else {
			//flash.Error("Invalid phone_number or password")
			//flash.Store(&c.Controller)
			output_table["ret_code"], output_table["message"] = -1, "Invalid phone_number or password"
		}
	} else {
		//flash.Error("Invalid json format")
		//flash.Store(&c.Controller)
		output_table["ret_code"], output_table["message"] = -1, "Invalid json data"
	}
	c.Data["json"] = output_table
	c.ServeJSON()
}

// Post: logout
func (c *UserController) Logout() {
	table := make(map[string]interface{})
	beego.ReadFromRequest(&c.Controller)
	token := c.GetString("token")
	fmt.Println("logout:", token)
	if hasRows, user := models.GetUserByToken(token); hasRows {
		//if err, user := filters.IsLogin(c.Ctx); err == true {
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

// Get: Get user information
func (c *UserController) GetUserInfo() {
	token := c.Input().Get("token")
	ret_data := make(map[string]interface{})
	type UserInfoOutput struct {
		Ret_code int                    `json:"ret_code"`
		Message  string                 `json:"message"`
		Data     map[string]interface{} `json:"data"`
	}

	// check token
	exist, user := models.GetUserByToken(token)
	if !exist {
		ret := UserInfoOutput{
			Ret_code: -2,
			Message:  "Invalid token",
			Data:     ret_data,
		}
		c.Data["json"] = &ret
		c.ServeJSON()
	} else {
		ret_data["id"] = user.Id
		ret_data["phone_number"] = user.PhoneNumber
		ret_data["nickname"] = user.Nickname
		ret_data["avatar"] = user.Avatar
		ret := UserInfoOutput{
			Ret_code: 0,
			Message:  "",
			Data:     ret_data,
		}
		c.Data["json"] = &ret
		c.ServeJSON()
	}
}
