package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"mini_proj/src/livingserver/utils"
	"strconv"
)

type Post struct {
	Id       	  string `orm:"pk;auto"`
	User          *User `orm:"rel(fk)"`
	Section       *Section `orm:"rel(fk)"`
	//Reply	      *Reply `orm:"reverse(many)"` // todo: 是否需要逆对应
	Content       string `orm:"type(text)"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)"`
	ViewCount	  uint64 `orm:"default(0)"`
	LikeCount     uint64 `orm:"default(0)"`
	ReplyCount    uint64 `orm:"default(0)"`
	StrongLevel	  uint	`orm:"default(0)"`
	Labels		  []string
	Visible		  bool
}

func AddPost(post *Post) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(post)
	return id
}

func FindPostById(id int) Post {
	o := orm.NewOrm()
	var post Post
	o.QueryTable(post).RelatedSel().Filter("Id", id).One(&post)
	return post
}

func PageTopic(p int, size int, section *Section) utils.Page {
	o := orm.NewOrm()
	var post Post
	var list []Post
	qs := o.QueryTable(post)
	if section.Id > 0 {
		qs = qs.Filter("Section", section)
	}
	count, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("-InTime").Limit(size).Offset((p - 1) * size).All(&list)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(c, p, size, list)
}

func IncrViewCount(post *Post) {
	o := orm.NewOrm()
	post.ViewCount = post.ViewCount + 1
	o.Update(post, "ViewCount")
}

func IncrLikeCount(post *Post) {
	o := orm.NewOrm()
	post.LikeCount = post.LikeCount + 1
	o.Update(post, "LikeCount")
}

func ReduceLikeCount(post *Post) {
	o := orm.NewOrm()
	post.LikeCount = post.LikeCount - 1
	o.Update(post, "LikeCount")
}

func IncrReplyCount(post *Post) {
	o := orm.NewOrm()
	post.ReplyCount = post.ReplyCount + 1
	o.Update(post, "ReplyCount")
}

func ReduceReplyCount(post *Post) {
	o := orm.NewOrm()
	post.ReplyCount = post.ReplyCount - 1
	o.Update(post, "ReplyCount")
}

func FindPostByUser(user *User, limit int) []*Post {
	o := orm.NewOrm()
	var post Post
	var posts []*Post
	o.QueryTable(post).RelatedSel().Filter("User", user).OrderBy("-CreateTime").Limit(limit).All(&posts)
	return posts
}

//func UpdatePost(post *Post) {
//	o := orm.NewOrm()
//	o.Update(post)
//}

func DeletePost(post *Post) {
	o := orm.NewOrm()
	o.Delete(post)
}

func DeletePostByUser(user *User) {
	o := orm.NewOrm()
	o.Raw("delete from post where user_id = ?", user.Id).Exec()
}