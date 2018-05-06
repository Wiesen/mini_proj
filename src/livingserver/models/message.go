package models

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Message struct {
	Id         int       `orm:"column(id);auto"`
	Poster     *User     `orm:"column(poster);rel(fk)" description:"发布人id"`
	Owner      *User     `orm:"column(owner);rel(fk)" description:"owner id"`
	Emotion    *Emotion  `orm:"column(emotion_id);rel(fk)" description:"心情ID"`
	Content    string    `orm:"column(content);size(256)" description:"评论内容"`
	TypeId     int       `orm:"column(type_id)"`
	CreateTime time.Time `orm:"column(create_time);type(datetime)" description:"时间"`
	//Like  *Like  		`orm:"column(like_id);rel(fk)" description:"ID"`
	//Comment  *Comment  	`orm:"column(comment_id);rel(fk)" description:"评论ID"`
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
	logs.Debug("add message[%+v] successful", *m)

	return
}

func GetMessageByUser(uid, pageNo int) (bool, []*Message) {
	o := orm.NewOrm()
	qs := o.QueryTable("message")
	var messages []*Message
	num, err := qs.Filter("owner", uid).OrderBy("-create_time").
		Limit(PAGE_SIZE, pageNo*PAGE_SIZE).All(&messages)
	// fmt.Println("Number of records retrieved in database:", num)
	logs.Debug("Number of records retrieved in database:", num)
	return (err != nil && err != orm.ErrNoRows), messages
}
