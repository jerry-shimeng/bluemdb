package bmnetwork

import (
	"bluem/bmcache"
	"bufio"
	"fmt"
	"net"
)

var (
	ip = "127.0.0.1"
)

type BmServer struct {
}

func init() {}

func (this *BmServer) Run(port int,ip string) {
	fmt.Println("listener server is running")
	go start(port,ip)
}

func start(port int,ip string) {
	var tcpAddr *net.TCPAddr
	serverAddr := fmt.Sprintf("%s:%d", ip, port)
	tcpAddr, _ = net.ResolveTCPAddr("tcp", serverAddr)
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	for {
		tcpConn, err := tcpListener.AcceptTCP()

		if err != nil {
			fmt.Println(err.Error())
		}

//		go tcpPipe(tcpConn)
		go tcpPipeChannel(tcpConn)
	}
}

//及时处理的模式
func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	fmt.Println("connected:", ipStr)
	defer func() {
		fmt.Println(" disconnected:", ipStr)
	}()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		fmt.Print("*")
		if err != nil {
			return
		}
		fmt.Printf("%s", (message))

		r := bmcache.ExecCommand(message)

		msg := processBack(r)
		b := []byte(msg)
		conn.Write(b)
	}
}
//改写上面的方法  用通道的模式
func tcpPipeChannel(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	fmt.Println("connected:", ipStr)
	defer func() {
		fmt.Println(" disconnected:", ipStr)
	}()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		//拿到数据
		if err != nil {
			return
		}
		//加入通道
		temp := bmcache.CommandModel{message,conn,execCallBack}
		bmcache.CommandChan <- temp
	}
}


//执行完毕后回调函数
func execCallBack(conn *net.TCPConn,s string){
	conn.Write([]byte(s))
}

func processBack(s string)string{
	return s+"\n"
}


