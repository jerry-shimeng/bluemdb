package bmcmdparam
import (
	"flag"
)

type CmdParamModel struct {
	Port int
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

	flag.Parse()

	CmdParamsObject.Port = *port
	CmdParamsObject.ConfigPath = *confPath

}