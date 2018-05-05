package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id         int       `orm:"column(id);auto"`
	Emotion  *Emotion  `orm:"column(emotion_id);rel(fk)" description:"心情ID"`
	Content    string    `orm:"column(content);size(256)" description:"评论内容"`
	Poster     *User     `orm:"column(poster);rel(fk)" description:"发布人id"`
	CreateTime time.Time `orm:"column(create_time);type(datetime)" description:"时间"`
	Rspto      int       `orm:"column(rspto);null" description:"被回复人id"`
}

func (t *Comment) TableName() string {
	return "comment"
}

func init() {
	orm.RegisterModel(new(Comment))
}

// AddComment insert a new Comment into database and returns
// last inserted Id on success.
// func AddComment(m *Comment) (id int64, err error) {
// 	o := orm.NewOrm()
// 	id, err = o.Insert(m)
// 	return
// }
func AddComment(m *Comment) (id int64, err error) {
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
	if err != nil {
		return
	}

	v := &Emotion{ Id : m.Emotion.Id}
	err = o.Read(v)
	if err != nil {
		return
	}

	v.CommentCnt++
	_, err = o.Update(v, "comment_cnt")

	return
}

// GetCommentById retrieves Comment by Id. Returns error if
// Id doesn't exist
func GetCommentById(id int) (v *Comment, err error) {
	o := orm.NewOrm()
	v = &Comment{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}


// DeleteComment deletes Comment by Id and returns error if
// the record to be deleted doesn't exist
func DeleteComment(id int) (err error) {
	o := orm.NewOrm()
	v := Comment{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Comment{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// ------------- following is added by wiesenyang -------------
func GetCommentByUser(uid, pageNo int) (bool, []*Comment) {
	o := orm.NewOrm()
	qs := o.QueryTable("comment")
	var comments []*Comment
	num, err := qs.Filter("poster", uid).OrderBy("-create_time").
		Limit(PAGE_SIZE, pageNo*PAGE_SIZE).All(&comments)
	fmt.Println("Number of records retrieved in database:", num)
	return (err != nil && err != orm.ErrNoRows), comments
}

func GetCommentByEmotion(emotionId, pageNo int) (bool, []*Comment) {
	o := orm.NewOrm()
	qs := o.QueryTable("comment")
	var comments []*Comment
	num, err := qs.Filter("emotion_id", emotionId).OrderBy("-create_time").
		Limit(PAGE_SIZE, pageNo*PAGE_SIZE).All(&comments)
	fmt.Println("Number of records retrieved in database:", num)
	return (err != nil && err != orm.ErrNoRows), comments
}

