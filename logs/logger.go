package logs

import (
	"fmt"
	"github.com/google/logger"
	"os"
	"time"
)

var Logger *logger.Logger

func InitLogger() {
	//制定是否控制台打印 默认为true
	fileName := fmt.Sprintf("../logs/" + time.Now().Format("2006_01_02") + ".log")

	lf, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		panic("Failed to open log file")
	}

	Logger = logger.Init("LoggerExample", true, false, lf)
}

//func Debug(format string, a ...interface{}) {
//	logger.Debug(fmt.Sprintf(format, a...))
//}
//func Info(format string, a ...interface{}) {
//	logger.Info(fmt.Sprintf(format, a...))
//}
//func Warn(format string, a ...interface{}) {
//	logger.Warn(fmt.Sprintf(format, a...))
//}
//func Error(format string, a ...interface{}) {
//	logger.Error(fmt.Sprintf(format, a...))
//}
//func Fatal(format string, a ...interface{}) {
//	logger.Fatal(fmt.Sprintf(format, a...))
//}
