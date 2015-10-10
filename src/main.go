package main

import (
	"bluem/bmcache"
	"bluem/bmnetwork"
	"fmt"
	"os"
	"os/signal"
	"runtime"

	"bluem/bmlog"

	"bluem/bmconfig"
	"bluem/bmcmdparam"
)


func main() {
	//使用多核cpu
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("bluemdb is starting...")

	bmcmdparam.CmdParamInit()
	//初始化配置
	bmconfig.ConfigInit(&bmcmdparam.CmdParamsObject)
	bmlog.LogInit()

	server := bmnetwork.BmServer{}

	port:= bmconfig.BmConfigObject.Port
	ip := bmconfig.BmConfigObject.Ip
	//启动监听
	server.Run(port,ip)

	//启动数据管理程序
	bmcache.CacheInit()

	fmt.Println("bluemdb is started ,host is ",ip,":",port)

	//程序正常退出捕获 ctrl+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	onExit()

}

//程序正常退出
func onExit() {
	fmt.Println("server is stoped \nbye!")
}
