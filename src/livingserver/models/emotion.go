package models

import (
	// "errors"
	"fmt"
	// "reflect"
	// "strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Emotion struct {
	Id         int       `orm:"column(id);auto"`
	Content    string    `orm:"column(content);size(256)" description:"心情内容"`
	LabelId    *Label    `orm:"column(label_id);rel(fk)" description:"心情标签ID，需存在标签表中"`
	Strong     int8      `orm:"column(strong)" description:"强度"`
	CreateTime time.Time `orm:"column(create_time);type(datetime)" description:"创建时间"`
	Visiable   int8      `orm:"column(visiable)" description:"1. 个人可见；2. 社区可见"`
	Poster     *User     `orm:"column(poster);rel(fk)" description:"发布人id"`
	CommentCnt uint      `orm:"column(comment_cnt)"`
	LikeCnt    uint      `orm:"column(like_cnt)"`
}

func (t *Emotion) TableName() string {
	return "emotion"
}

func init() {
	orm.RegisterModel(new(Emotion))
}

// AddEmotion insert a new Emotion into database and returns
// last inserted Id on success.
func AddEmotion(m *Emotion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEmotionById retrieves Emotion by Id. Returns error if
// Id doesn't exist
func GetEmotionById(id int) (v *Emotion, err error) {
	o := orm.NewOrm()
	v = &Emotion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEmotion retrieves all Emotion matches certain condition. Returns empty list if
// no records exist
// func GetAllEmotion(query map[string]string, fields []string, sortby []string, order []string,
// 	offset int64, limit int64) (ml []interface{}, err error) {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable(new(Emotion))
// 	// query k=v
// 	for k, v := range query {
// 		// rewrite dot-notation to Object__Attribute
// 		k = strings.Replace(k, ".", "__", -1)
// 		if strings.Contains(k, "isnull") {
// 			qs = qs.Filter(k, (v == "true" || v == "1"))
// 		} else {
// 			qs = qs.Filter(k, v)
// 		}
// 	}
// 	// order by:
// 	var sortFields []string
// 	if len(sortby) != 0 {
// 		if len(sortby) == len(order) {
// 			// 1) for each sort field, there is an associated order
// 			for i, v := range sortby {
// 				orderby := ""
// 				if order[i] == "desc" {
// 					orderby = "-" + v
// 				} else if order[i] == "asc" {
// 					orderby = v
// 				} else {
// 					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
// 				}
// 				sortFields = append(sortFields, orderby)
// 			}
// 			qs = qs.OrderBy(sortFields...)
// 		} else if len(sortby) != len(order) && len(order) == 1 {
// 			// 2) there is exactly one order, all the sorted fields will be sorted by this order
// 			for _, v := range sortby {
// 				orderby := ""
// 				if order[0] == "desc" {
// 					orderby = "-" + v
// 				} else if order[0] == "asc" {
// 					orderby = v
// 				} else {
// 					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
// 				}
// 				sortFields = append(sortFields, orderby)
// 			}
// 		} else if len(sortby) != len(order) && len(order) != 1 {
// 			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
// 		}
// 	} else {
// 		if len(order) != 0 {
// 			return nil, errors.New("Error: unused 'order' fields")
// 		}
// 	}

// 	var l []Emotion
// 	qs = qs.OrderBy(sortFields...)
// 	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
// 		if len(fields) == 0 {
// 			for _, v := range l {
// 				ml = append(ml, v)
// 			}
// 		} else {
// 			// trim unused fields
// 			for _, v := range l {
// 				m := make(map[string]interface{})
// 				val := reflect.ValueOf(v)
// 				for _, fname := range fields {
// 					m[fname] = val.FieldByName(fname).Interface()
// 				}
// 				ml = append(ml, m)
// 			}
// 		}
// 		return ml, nil
// 	}
// 	return nil, err
// }

// UpdateEmotion updates Emotion by Id and returns error if
// the record to be updated doesn't exist
func UpdateEmotionById(m *Emotion) (err error) {
	o := orm.NewOrm()
	v := Emotion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEmotion deletes Emotion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEmotion(id int) (err error) {
	o := orm.NewOrm()
	v := Emotion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Emotion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// ------------- following is added by yyff -------------
func GetEmotionByUser(uid, pageNo int) (bool, []*Emotion) {
	o := orm.NewOrm()
	qs := o.QueryTable("emotion")
	var emotions []*Emotion
	num, err := qs.Filter("poster", uid).OrderBy("-create_time").
		Limit(PAGE_SIZE, pageNo * PAGE_SIZE).All(&emotions)
	fmt.Println("Number of records retrieved in database:", num)
	return (err != nil && err != orm.ErrNoRows), emotions
}

func GetAllEmotion(pageNo int) (bool, []*Emotion) {
	o := orm.NewOrm()
	qs := o.QueryTable("emotion")
	var emotions []*Emotion
	num, err := qs.Filter("visiable", 1).OrderBy("-create_time").
		Limit(PAGE_SIZE, pageNo * PAGE_SIZE).All(&emotions)
	fmt.Println("Number of records retrieved in database:", num)
	return (err != nil && err != orm.ErrNoRows), emotions
}

func GetEmotionByLabel(labelId, pageNo int) (bool, []*Emotion) {
	o := orm.NewOrm()
	qs := o.QueryTable("emotion")
	var emotions []*Emotion
	num, err := qs.Filter("visiable", 1).Filter("label_id", labelId).OrderBy("-create_time").
		Limit(PAGE_SIZE, pageNo * PAGE_SIZE).All(&emotions)
	fmt.Println("Number of records retrieved in database:", num)
	return (err != nil && err != orm.ErrNoRows), emotions
}