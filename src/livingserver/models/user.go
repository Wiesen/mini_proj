package models

import (
	"github.com/astaxie/beego/orm"
)


type User struct {
	Id      string	`orm:"pk;unique"`		// phone_number
	QQ 		string	`orm:"unique"`
	Username string	`orm:"unique"`
	Password string
	Token    string `orm:"unique"`
	Avatar   uint	`orm:"default(0)"`
	//Post    []*Post `orm:"reverse(many)"`
}

func AddUser(user *User) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(user)
	return id
}

func FindUserById(Id string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Id", Id).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByToken(token string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Token", token).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByQQ(QQ string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("QQ", QQ).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByUserName(username string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", username).One(&user)
	return err != orm.ErrNoRows, user
}

func UpdateUser(user *User) (int64, error) {
	o := orm.NewOrm()
	return o.Update(user)
}

func LoginByID(Id string, password string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Id", Id).Filter("Password", password).One(&user)
	return err != orm.ErrNoRows, user
}

func LoginByQQ(QQ string, password string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("QQ", QQ).Filter("Password", password).One(&user)
	return err != orm.ErrNoRows, user
}

func DeleteUser(user *User) {
	o := orm.NewOrm()
	o.Delete(user)
}
