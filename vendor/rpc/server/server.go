package server

import (
	"github.com/astaxie/beego"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/hprose/hprose-golang/rpc/websocket"
)

var Server = websocket.NewWebSocketService()

type Event struct{}

// 监听上线(推送)
func (Event) OnSubscribe(topic string, id string, service rpc.Service) {
	beego.Debug("client " + id + " online subscribe topic is: " + topic)
}

// 监听离线(推送)
func (Event) OnUnsubscribe(topic string, id string, service rpc.Service) {
	beego.Debug("client " + id + " offline unsubscribe topic is: " + topic)
}

// 推送
func Push(topic string, result interface{}, id ...string) {
	Server.Push(topic, result, id...)
}

// 广播
func Broadcast(topic string, result interface{}, callback func([]string)) {
	Server.Broadcast(topic, result, callback)
}

// 多播
func Multicast(topic string, ids []string, result interface{}, callback func([]string)) {
	Server.Multicast(topic, ids, result, callback)
}

// 单播
func Unicast(topic string, id string, result interface{}, callback func(bool)) {
	Server.Unicast(topic, id, result, callback)
}
