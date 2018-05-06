package main

import (
	_ "livingserver/redis_client"
	_ "livingserver/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"livingserver.log"}`)
}

func main() {
	beego.Run()
}
