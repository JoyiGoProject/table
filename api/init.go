package controllers

import (
	m "table/models"

	"github.com/astaxie/beego"
)

type Response struct {
	Status bool
	Info   string
	Data   interface{}
}

func Init(uid string) (res Response) {
	beego.Info("远程数据", uid)
	res.Status = true
	res.Info = "数据接收成功"
	/*	u := m.GetUid(uid)
		beego.Info(u)
		if u != uid {
			res.Info = "数据错误"
			return res
		}*/
	ud := m.SetTableByUid(uid)
	beego.Info("建立表返回数据：", ud)
	if ud != uid {
		res.Info = "数据错误"
		return res
	}
	return res
}
