package filters

import (
	"livingserver/models"

	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

func IsLogin(ctx *context.Context) (bool, models.User) {
	token := ctx.Input.Param(":token")
	// fmt.Println("get token: ", token)
	logs.Debug("get token: ", token)
	err, user := models.GetUserByToken(token)
	return err, user
}

var FilterUser = func(ctx *context.Context) {
	ok, _ := IsLogin(ctx)
	if !ok {
		ctx.Redirect(302, "/login")
	}
}
