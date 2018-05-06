package controllers

import (
	"config"
	"fmt"
	"log"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var conf struct {
	Database struct {
		User     string `default:"root"`
		Password string `default:"123456"`
		Address  string `default:"127.0.0.1:3306"`
		Dbname   string `default:"livingdb"`
	}
}

func init() {
	config.Parse(&conf)
	sqlConfig := fmt.Sprintf("%s:%s@tcp(%s)/%s", conf.Database.User,
		conf.Database.Password, conf.Database.Address, conf.Database.Dbname)

	err := orm.RegisterDataBase("default", "mysql", sqlConfig)
	if err != nil {
		log.Fatal("init database failed")
	}
	log.Println("Init database successful")
}

type PostLikeReq struct {
	EmotionID int `json:"emotion_id"`
}

type PostEmotionReq struct {
	Content  string `json:"content"`
	LabelID  int    `json:"label_id"`
	Strong   int8   `json:"strong"`
	Visiable int8   `json:"visiable"`
}

type PostCommentReq struct {
	Comment   string `json:"comment"`
	EmotionID int    `json:"emotion_id"`
	RspTo     int    `json:"rspto"`
}

type CommonRsp struct {
	RetCode int           `json:"ret_code"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data"`
}
