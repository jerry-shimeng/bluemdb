package bmcache

import (
	"bluem/bmwatch"
	"errors"
	"strings"
)

const (
	ParamErr        = "Parameter error"
	CommandNotFound = " Command Not Found"
	//CommandArray    = []string{"add", "set", "del"}
)


var commandMap map[string]func(string, string) string

func CommandInit() {
	commandMap = make(map[string]func(string, string) string, 0)
	commandMap["add"] = Add
	commandMap["set"] = Set
	commandMap["get"] = Get
	commandMap["del"] = Del
	commandMap["append"] =Append
	commandMap["getset"] = GetSet
	commandMap["err"] = bmwatch.ErrorInfo
}

//执行接受的命令
func ExecCommand(content string) string {

	content = strings.Replace(content, "\n", "", -1)
	//json 解析
	cmd, key, value, err := convertConmmand(content)
	if err != nil {
		return ParamErr
	}

	//cmd 转换为小写
	cmd = strings.ToLower(cmd)

	//验证命令是否存在
	b := existCommand(cmd)

	if b == true {
		//执行
		s := commandMap[cmd](key, value)
		return s

	} else {
		return CommandNotFound
	}

	return ParamErr
}
//转换命令
func convertConmmand(s string) (string, string, string, error) {
	if len(s) == 0 {
		return "", "", "", errors.New("param is null")
	}
	arr := strings.Split(s, " ")
	arr = removeArraySpace(arr)
	if len(arr) == 2 || len(arr) == 3 {
		//处理
		cmd := arr[0]
		key := arr[1]
		value := ""
		if len(arr) == 3 {
			value = arr[2]
		}
		return cmd, key, value, nil
	}
	return "", "", "", errors.New("param is null")
}
//判断该命令是否存在
func existCommand(key string) bool {
	r := commandMap[key]
	if r != nil {
		return true
	} else {
		return false
	}
}
//移除多余的空格
func removeArraySpace(arr []string) []string {

	s := make([]string, 0)
	for _, v := range arr {
		if v != "" && v != " " {
			s = append(s, v)
		}
	}
	//fmt.Println(s)
	return s
}

