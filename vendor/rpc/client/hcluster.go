package client

import (
	"fmt"

	"github.com/astaxie/beego"

	"github.com/hprose/hprose-golang/rpc"
	_ "github.com/hprose/hprose-golang/rpc/websocket"
)

type ErrorEvent struct{}

func (e *ErrorEvent) OnError(name string, err error) {
	fmt.Println(name, err.Error())
}

type HresFunc struct {
	PushDeviceList func(servercode string, number, typeid, userid, groupid, grouplimit, userlimit int64) (status bool, info string, data *[]int64)
	// 检查服务器状态
	ServerStatus func() error
}

func UserHres(serverip, serverport string) (clientHres *HresFunc, err error) {
	url := "ws://" + serverip + ":" + serverport + "/rpc"
	beego.Debug("request url", url)
	client := rpc.NewClient(url)
	client.SetEvent(ErrorEvent{})
	client.UseService(&clientHres)
	err = clientHres.ServerStatus()
	return clientHres, err
}
