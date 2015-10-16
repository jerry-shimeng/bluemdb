package bmmodel

import (
	"strings"
)

type CmdModel struct {
	Cmd   string
	Key   string
	Value string
}

//字符串转换为命令模型
func ConvertStringToCmd(s string) *CmdModel {
	model := jsonToCmdModel(&s)
	return model
}

//获取下个命令内容
func jsonToCmdModel(s *string) *CmdModel {
	content := strings.Replace(*s, "\n", "", -1)
	v := JsonConvertToObject(content)
	return v
}


