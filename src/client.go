package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"flag"
	"errors"
	"strings"
	"encoding/json"
)

type CmdModel struct {
	Cmd   string
	Key   string
	Value string
}

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

		s,err := getCmdString(s)
		if err!=nil {
			fmt.Println(err.Error())
		}else {
			conn.Write([]byte(s + "\n"))
		}
	}
}

func getCmdString(s string) (string,error){
	cmd,key,value,err := convertConmmand(s)
	if err !=nil {
		return "",err
	}
	model :=CmdModel{cmd,key,value}
	str,err := json.Marshal(model)
	if err ==nil {
		temp := string(str)
		return temp,nil
	}else {
		return "",err
	}
}


//转换命令
func convertConmmand(s string) (string, string, string, error) {
	if len(s) == 0 {
		return "", "", "", errors.New("param is null")
	}
	arr := strings.Split(s, " ")
	//arr = removeArraySpace(arr)
	fmt.Println(arr)
	var cmd, key,value ="","",""

	for i:=0;i<len(arr) ;i++  {
			if cmd ==""{
				if (arr[i] != ""){
					cmd = arr[i]
				}
			}else if key  == ""{
				if (arr[i] != "") {
					key = arr[i]
				}
			}else {
				if arr[i]=="" {
					value += " "
				}else {
					value+=arr[i]
				}
			}
	}

//	if len(arr) == 2 || len(arr) >= 3 {
//		//处理
//		cmd := arr[0]
//		key := arr[1]
//		value := ""
//		if len(arr) == 3 {
//			for i :=2; i< len(arr) ;i++ {
//				value = value +arr[i]
//			}
//		}
//		return cmd, key, value, nil
//	}
	return cmd, key, value, nil
}
//移除多余的空格
func removeArraySpace(arr []string) []string {

	s := make([]string, 0)
	for _, v := range arr {
		if v != "" && v != " " {
			s = append(s, v)
		}
	}
	return s
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
	ip := flag.String("ip","192.168.155.239","bluemdb server ip address")

	flag.Parse()

	serverHost = fmt.Sprintf("%s:%d",*ip,*port)

}