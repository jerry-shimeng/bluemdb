package bmlog

import (
	"github.com/astaxie/beego/logs"
	"bluem/bmwatch"
)

var bmlog *logs.BeeLogger

func LogInit() {
	bmlog = logs.NewLogger(10000)
	bmlog.SetLogger("file", `{"filename":"test.log"}`)
}

func Debug(str string) {
	bmlog.Debug(str)
}

func Info(str string) {
	bmlog.Info("%s",str)
}

func Error(err error) {

	bmwatch.ErrorWatchLog(&err)

	bmlog.Error(err.Error())
}

func ErrorMsg(str string, err error) {
	bmwatch.ErrorWatchLog(&err)

	bmlog.Error("%s:%d",str, err)
}
