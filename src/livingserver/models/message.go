package models

import "github.com/astaxie/beego/orm"

type Message struct {
	Id        int    	`orm:"column(id);auto"`
	Like  *Like  		`orm:"column(like_id);rel(fk)" description:"ID"`
	Comment  *Comment  	`orm:"column(comment_id);rel(fk)" description:"评论ID"`
}

func (t *Message) TableName() string {
	return "message"
}

func init() {
	orm.RegisterModel(new(Message))
}

// AddComment insert a new Comment into database and returns
// last inserted Id on success.
// func AddComment(m *Comment) (id int64, err error) {
// 	o := orm.NewOrm()
// 	id, err = o.Insert(m)
// 	return
// }
func AddMessage(m *Message) (id int64, err error) {
	o := orm.NewOrm()

	defer func() {
		if err != nil {
			o.Rollback()
		} else {
			o.Commit()
		}
	}()

	// 事务
	err = o.Begin()
	id, err = o.Insert(m)

	return
}
