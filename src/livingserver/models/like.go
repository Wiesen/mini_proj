package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Like struct {
	Id         int       `orm:"column(id);auto"`
	EmotionId  *Emotion  `orm:"column(emotion_id);rel(fk)" description:"心情id"`
	Poster     *User     `orm:"column(poster);rel(fk)" description:"发布人id"`
	CreateTime time.Time `orm:"column(create_time);type(datetime)" description:"时间"`
}

func (t *Like) TableName() string {
	return "like"
}

func init() {
	orm.RegisterModel(new(Like))
}

// GetLikeById retrieves Like by Id. Returns error if
// Id doesn't exist
func GetLikeById(id int) (v *Like, err error) {
	o := orm.NewOrm()
	v = &Like{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllLike retrieves all Like matches certain condition. Returns empty list if
// no records exist
func GetAllLike(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Like))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Like
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateLike updates Like by Id and returns error if
// the record to be updated doesn't exist
func UpdateLikeById(m *Like) (err error) {
	o := orm.NewOrm()
	v := Like{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLike deletes Like by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLike(id int) (err error) {
	o := orm.NewOrm()
	v := Like{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Like{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}


// following is added by yyff
func GetLikeByUser(uid int) (bool, []*Like) {
	o := orm.NewOrm()
	qs := o.QueryTable("like")
	var likes []*Like
	num, err := qs.Filter("poster", uid).All(&likes)
	if err != nil {
		fmt.Println("query table failed, err info: %+v", err)
	}
	fmt.Println("Number of records retrieved in database:", num)
	return (err != nil && err != orm.ErrNoRows), likes
}

// AddLike insert a new Like into database and returns
// last inserted Id on success.
// func AddLike(m *Like) (id int64, err error) {
// 	o := orm.NewOrm()
// 	id, err = o.Insert(m)
// 	return
// }
func AddLike(m *Like) (id int64, err error) {
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

	v := &Emotion{ Id : m.EmotionId.Id}
	err = o.Read(v)
	if err != nil {
		return
	}

	v.LikeCnt++
	_, err = o.Update(v)
	return

}
