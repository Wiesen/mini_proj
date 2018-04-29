package models

import (
	"image"
	"github.com/astaxie/beego/orm"
)


type User struct {
	uId      uint64		// phone_number, unique
	QQNumber uint64		// qq_number, unique
	Username string		// nickname, unique
	Password string
	Avatar   image.Image
}

func AddUser(user *User) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(user)
	return id
}

func FindUserById(uId uint64) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("uId", uId).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByQQ(QQ uint64) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("QQNumber", QQ).One(&user)
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

func LoginByName(username string, password string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", username).Filter("Password", password).One(&user)
	return err != orm.ErrNoRows, user
}

func LoginByUID(uId uint64, password string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("uId", uId).Filter("Password", password).One(&user)
	return err != orm.ErrNoRows, user
}

func LoginByQQ(QQ uint64, password string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("QQNumber", QQ).Filter("Password", password).One(&user)
	return err != orm.ErrNoRows, user
}

func DeleteUser(user *User) {
	o := orm.NewOrm()
	o.Delete(user)
}
