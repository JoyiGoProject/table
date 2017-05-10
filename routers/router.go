package routers

import (
	api "table/api"
	"table/models"
	"table/rpc"

	"github.com/astaxie/beego"
)

func init() {
	initialize()
	Rpcserver()
}

func initialize() {
	models.InitDB()
}

func Rpcserver() {
	rpc.Server.Event = rpc.Event{}
	rpc.Server.AddFunction("InitUser", api.Init)
	beego.Handler("/rpc", rpc.Server)
}
