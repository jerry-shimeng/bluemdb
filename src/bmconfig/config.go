//配置模块
package bmconfig
import (
	"github.com/astaxie/beego/config"
	"fmt"
	"bluem/bmcmdparam"
	"net"
	"strings"
)

var BmConfigObject BmConfigModel

//配置模型
type BmConfigModel struct {
	Port int
	Ip string
}

func ConfigInit(param *bmcmdparam.CmdParamModel){
	configModelDefault()

	if len(param.ConfigPath)>0{
		//配置文件
		bm, err := config.NewConfig("ini", param.ConfigPath)
		if err != nil {
			fmt.Println("config init error :",err.Error())
			fmt.Println("config default init ")
		}else{
			setConfigModel(bm)
		}
	}


	if param.Port > 0{
		BmConfigObject.Port = param.Port
	}
}

//初始化默认的配置模型
func configModelDefault(){
	BmConfigObject = BmConfigModel{Port:8090}
	BmConfigObject.Ip =getIp()
}
func getIp() string{
	conn, err := net.Dial("udp", "g.cn:80")
	if err != nil {
		fmt.Println("get local ip address error ",err.Error())
		return "127.0.0.1"
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}
func setConfigModel(bm config.ConfigContainer){

	//读取端口号
	BmConfigObject.Port = bm.DefaultInt("port",8090)
}
