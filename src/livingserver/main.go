package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"mini_proj/src/livingserver/models"
)

func initDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123@tcp(127.0.0.1:3306)/database?charset=utf8")
	orm.RegisterModel(
		new(models.User),
		new(models.Section),
		new(models.Post))
	orm.RunSyncdb("default", false, true)
}

func main() {
	initDB()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
