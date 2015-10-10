package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"flag"
)

var continuechan chan bool
var serverHost string
func main() {
	cmd()
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp4", serverHost		)

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	fmt.Println("connected!")

	go onMessageRecived(conn)
	var s string
	continuechan = make(chan bool, 1)

	continuechan <- true

	for {
		<-continuechan
		fmt.Print("bule-clent: ")
		s = consoleReadline()
		conn.Write([]byte(s + "\n"))
	}
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Print(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
		continuechan <- true
	}
}

func consoleReadline()string{
	reader := bufio.NewReader(os.Stdin)
	b,_,_ := reader.ReadLine()
	return string(b)
}


func cmd(){
	port := flag.Int("port", 8090, "bluemdb server port")
	ip := flag.String("ip","127.0.0.1","bluemdb server ip address")

	flag.Parse()

	serverHost = fmt.Sprintf("%s:%d",*ip,*port)

}