package main

import (
	_ "livingserver/controllers"
	_ "livingserver/redis_client"
	_ "livingserver/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
