package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id          int    `orm:"column(id);auto"`
	QqNumber    string `orm:"column(qq_number);size(64);null;unique" description:"qq号"`
	PhoneNumber string `orm:"column(phone_number);size(64);unique" description:"手机号"`
	Password    string `orm:"column(password);size(64)"`
	Nickname    string `orm:"column(nickname);size(64);unique" description:"昵称"`
	Token       string `orm:"column(token);size(64);null" description:"登录token"`
	Avatar      string `orm:"column(avatar);size(256);null" description:"头像"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return err
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// the following is added by Wiesenyang
func GetUserByToken(token string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Token", token).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByUsername(username string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", username).One(&user)
	return err != orm.ErrNoRows, user
}

func Login(phoneNumber string, password string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("PhoneNumber", phoneNumber).Filter("Password", password).One(&user)
	return err != orm.ErrNoRows, user

}

func GenerateUsername(username string) string {
	o := orm.NewOrm()
	var count int
	search := username + "[0-9]*$"
	o.Raw("select count(*) as Count from user where nickname REGEXP ?", search).QueryRow(&count)
	// fmt.Println("GenerateUsername:", count)
	logs.Debug("GenerateUsername:", count)
	if int(count) == 0 {
		return username
	} else {
		return username + strconv.Itoa(int(count))
	}
}
