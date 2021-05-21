// Author: coolliu
// Date: 2021/5/21

package log

import (
	"io"
	golog "log"
	"os"
)

type Log interface {
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

type DefaultLog struct {
}

func (l DefaultLog) Fatalf(format string, args ...interface{}) {
	golog.Fatalf(format, args...)
}
func (l DefaultLog) Fatal(args ...interface{}) {
	golog.Fatal(args...)
}
func (l DefaultLog) Errorf(format string, args ...interface{}) {
	golog.Printf("error|"+format, args...)
}
func (l DefaultLog) Error(args ...interface{}) {
	golog.Printf("error|", args...)
}

func (l DefaultLog) Warnf(format string, args ...interface{}) {
	golog.Printf("warn|", args...)
}
func (l DefaultLog) Warn(args ...interface{}) {
	golog.Printf("warn|", args...)
}
func (l DefaultLog) Infof(format string, args ...interface{}) {
	golog.Printf("info|"+format, args...)
}
func (l DefaultLog) Info(args ...interface{}) {
	golog.Printf("info|", args...)
}
func (l DefaultLog) Debugf(format string, args ...interface{}) {
	golog.Printf("debug|"+format, args...)
}
func (l DefaultLog) Debug(args ...interface{}) {
	golog.Printf("debug|", args...)
}
func (l DefaultLog) SetOutput(w io.Writer) {
	golog.SetOutput(w)
}

func NewDefaultLog() Log {
	log := DefaultLog{}
	log.SetOutput(os.Stdout)
	return log
}
