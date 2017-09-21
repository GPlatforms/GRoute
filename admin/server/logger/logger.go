package logger

import (
	"os"

	"github.com/happyEgg/wlog"
)

func ErrorDiary() *wlog.WLogger {
	ErrLogs := wlog.NewLogger(100)
	ErrLogs.Async()
	ErrLogs.EnableFuncCallDepth(true)

	//创建err_diary.log文件
	os.Mkdir("diary", os.ModePerm)
	ErrLogs.SetLogger("file", `{"filename":"./diary/err_diary.log"}`)

	return ErrLogs
}
