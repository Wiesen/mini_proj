package main

import (
	"os"

	_ "github.com/Wiesen/mini_proj/livingserver/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	sqlConfig := os.Getenv("LV_SQL")
	if 0 == len(sqlConfig) {
		// for local debug
		sqlConfig = "root:123456@tcp(127.0.0.1:3306)/livingdb"
	}
	orm.RegisterDataBase("default", "mysql", sqlConfig)
}

func main() {
	beego.BConfig.CopyRequestBody = true
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
