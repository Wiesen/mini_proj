package main

import (
	"log"
	"os"

	_ "livingserver/redis_client"
	_ "livingserver/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	sqlConfig := os.Getenv("LV_SQL")
	if 0 == len(sqlConfig) {
		// for local debug
		sqlConfig = "root:123456@tcp(127.0.0.1:3306)/livingdb"
	}
	err := orm.RegisterDataBase("default", "mysql", sqlConfig)
	if err != nil {
		log.Fatal("init database failed")
	}
	// logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":1000,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.SetLogger(logs.AdapterFile, `{"filename":"test.log"}`)
}

func main() {
	// redis_client.Init()
	beego.BConfig.CopyRequestBody = true
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
