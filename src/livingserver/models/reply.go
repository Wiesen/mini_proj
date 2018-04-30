package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Reply struct {
	Id     string `orm:"pk;auto"`
	Post   *Post  `orm:"rel(fk)"`
	User    *User `orm:"rel(fk)"`
	Content string `orm:"type(text);null"`
	CreateTime  time.Time `orm:"auto_now_add;type(datetime)"`
}

func AddReply(reply *Reply) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(reply)
	return id
}

func FindReplyById(id int) Reply {
	o := orm.NewOrm()
	var reply Reply
	o.QueryTable(reply).RelatedSel("Post").Filter("Id", id).One(&reply)
	return reply
}

func FindReplyByPost(post *Post) []*Reply {
	o := orm.NewOrm()
	var reply Reply
	var replies []*Reply
	o.QueryTable(reply).RelatedSel().Filter("Post", post).OrderBy("-Up", "-CreateTime").All(&replies)
	return replies
}

func FindReplyByUser(user *User, limit int) []*Reply {
	o := orm.NewOrm()
	var reply Reply
	var replies []*Reply
	o.QueryTable(reply).RelatedSel("Post", "User").Filter("User", user).OrderBy("-CreateTime").Limit(limit).All(&replies)
	return replies
}


func DeleteReply(reply *Reply) {
	o := orm.NewOrm()
	o.Delete(reply)
}

func DeleteReplyByPost(post *Post) {
	o := orm.NewOrm()
	var reply Reply
	var replies []Reply
	o.QueryTable(reply).Filter("Post", post).All(&replies)
	for _, reply := range replies {
		o.Delete(&reply)
	}
}

func DeleteReplyByUser(user *User) {
	o := orm.NewOrm()
	o.Raw("delete form reply where user_id = ?", user.Id).Exec()
}