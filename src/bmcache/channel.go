package bmcache
import (
	"net"
//	"fmt"
)


//命令通道
var	CommandChan chan CommandModel

type CommandModel struct  {
	CommandText string
	ConnPointer *net.TCPConn
	CallBack func(*net.TCPConn,string)
}

func ChannelInit(){
	CommandChan = make(chan CommandModel,500)
    go	ListenChannel()
}

//监视通道
func ListenChannel(){
	var obj CommandModel

	for{
		select {
		case obj = <-CommandChan:
			{
				//接受到请求，开始处理
				//1.执行命令
				r:=ExecCommand(obj.CommandText)
				//2.加上\n
				r = processBack(r)
				//3.回调函数
				obj.CallBack(obj.ConnPointer,r)
			}
		}
	}
}


func processBack(s string)string{
	return s+"\n"
}
