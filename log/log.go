// Author: coolliu
// Date: 2021/7/23

package log

type LogI interface {
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
	Warnf(format string, args ...interface{})
	Warn(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
}

func Fatalf(format string, args ...interface{}) {
	gLog.Fatalf(format, args...)
}
func Fatal(args ...interface{}) {
	gLog.Fatal(args...)
}
func Errorf(format string, args ...interface{}) {
	gLog.Errorf(format, args...)
}
func Error(args ...interface{}) {
	gLog.Error(args...)
}
func Warnf(format string, args ...interface{}) {
	gLog.Warnf(format, args...)
}
func Warn(args ...interface{}) {
	gLog.Warn(args...)
}
func Infof(format string, args ...interface{}) {
	gLog.Infof(format, args...)
}
func Info(args ...interface{}) {
	gLog.Info(args...)
}
func Debugf(format string, args ...interface{}) {
	gLog.Debugf(format, args...)
}
func Debug(args ...interface{}) {
	gLog.Debug(args...)
}

// 设置全局log
var gLog LogI

func init() {
	if gLog == nil {
		gLog = NewDefaultLog()
	}
}

func SetGlobalLog(log LogI) {
	gLog = log
}
