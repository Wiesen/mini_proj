package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"mini_proj/src/livingserver/models"
)

func initDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
	orm.RegisterModel(
		new(models.User))
	orm.RunSyncdb("default", false, true)
	//o := orm.NewOrm()
	//o.Using("default")
}

func main() {
	initDB()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
