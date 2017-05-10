package rpc

import (
	"github.com/astaxie/beego"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/hprose/hprose-golang/rpc/websocket"
)

var Server = websocket.NewWebSocketService()

type Event struct{}

// 监听上线
func (Event) OnSubscribe(topic string, id string, service rpc.Service) {
	beego.Debug("client " + id + " online subscribe topic is: " + topic)
}

// 监听离线
func (Event) OnUnsubscribe(topic string, id string, service rpc.Service) {
	beego.Debug("client " + id + " offline unsubscribe topic is: " + topic)
}

func Push(topic string, result interface{}, id ...string) {
	Server.Push(topic, result, id...)
}

func Broadcast(topic string, result interface{}, callback func([]string)) {
	Server.Broadcast(topic, result, callback)
}

func Multicast(topic string, ids []string, result interface{}, callback func([]string)) {
	Server.Multicast(topic, ids, result, callback)
}

func Unicast(topic string, id string, result interface{}, callback func(bool)) {
	Server.Unicast(topic, id, result, callback)
}
