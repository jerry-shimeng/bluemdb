package bmwatch

import "fmt"

var (
	errCount  = 0
	lastError *error
)

//错误记录
func ErrorWatchLog(err *error) {
	errCount++
	lastError = err
}

////获取错误的数量
//func GetErrorCount() int {
//	return errCount
//}
//
////最后一条错误信息
//func GetLastError() error {
//	return *lastError
//}

func ErrorInfo(key string, value string) string {
	var s string
	switch key {
	case "count":
		s = fmt.Sprintf("%d", errCount)
	case "last":
		s = fmt.Sprintf("%v", lastError)
	}
	return s
}
