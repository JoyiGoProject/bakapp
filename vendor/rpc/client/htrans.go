package client

import (
	"github.com/astaxie/beego"
	"github.com/hprose/hprose-golang/rpc"
	_ "github.com/hprose/hprose-golang/rpc/websocket"
	// "rpc/protocol"
)

type HdaemonFunc struct {
	// key 校验码 path 文件目录 name 项目名称 version 版本 argument 参数 enable 是否开启 autostart 是否自动启动 timedrestart 凌晨自启动 data 程序文件
	GetDeviceList    func(key, path, name, version, argument string, enable, autostart, timedrestart bool, data []byte) bool
	GetHdaemonStatus func() error
}

func UseHdaemon(serverip, serverport string) (clientHdaemon *HdaemonFunc, err error) {
	url := "ws://" + serverip + ":" + serverport
	beego.Debug("url", url)
	client := rpc.NewClient(url)
	client.UseService(&clientHdaemon)
	err = clientHdaemon.GetHdaemonStatus()
	return clientHdaemon, err
}
