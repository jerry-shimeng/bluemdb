package bmmodel

import (
	"encoding/json"
	"fmt"
)

//使用json
func JsonConvertToObject(s string) *CmdModel {
	var v CmdModel
	err := json.Unmarshal([]byte(s), &v)
	if err !=nil {
		fmt.Println(err.Error())
	}
	return &v
}

func JsonConvertToString(v *interface{}) string {

	b, _ := json.Marshal(*v)
	s := string(b)
	return s
}
