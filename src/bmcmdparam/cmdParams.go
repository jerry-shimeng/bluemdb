package bmcmdparam
import (
	"flag"
)

type CmdParamModel struct {
	Port int
	Host string
	ConfigPath string
}

var CmdParamsObject CmdParamModel

func CmdParamInit(){
	CmdParamsObject = CmdParamModel{}

	cmd()
}

func cmd(){
	port := flag.Int("port", 0, "http listen port")
	confPath := flag.String("conf","","bluemdb config file path")
	host := flag.String("host","","binding local host ip adress")
	flag.Parse()

	CmdParamsObject.Port = *port
	CmdParamsObject.Host = *host
	CmdParamsObject.ConfigPath = *confPath

}