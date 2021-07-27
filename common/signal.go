// Author: coolliu
// Date: 2021/5/21

package common

import (
	"github.com/liuping001/fastgo/log"
	"os"
	"os/signal"
	"syscall"
)

type OnSignal interface {
	OnExit()
	OnUser1()
	OnUser2()
}

type DefaultOnSignal struct {
}

func (a *DefaultOnSignal) OnExit() {
}
func (a *DefaultOnSignal) OnUser1() {

}
func (a *DefaultOnSignal) OnUser2() {

}

// 优雅退出
func RegisterSignal(onSignal OnSignal) {

	chanSignal := make(chan os.Signal)
	// 监听信号
	signal.Notify(chanSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)

	go func() {
		for s := range chanSignal {
			switch s {
			case syscall.SIGINT, syscall.SIGTERM:
				log.Infof("on recv quit signal")
				if onSignal != nil {
					onSignal.OnExit()
				}
			case syscall.SIGUSR1:
				log.Infof("on signal user")
				if onSignal != nil {
					onSignal.OnUser1()
				}
			case syscall.SIGUSR2:
				log.Infof("on signal user2")
				if onSignal != nil {
					onSignal.OnUser2()
				}
			default:
				log.Infof("other signal:%v", s)
			}
		}
	}()
}

type onExitSignal struct {
	DefaultOnSignal
	exitFunc func()
}

func (a *onExitSignal) OnExit() {
	if a.exitFunc != nil {
		a.exitFunc()
	}
}

func RegisterSignalExit(f func()) {
	if f == nil {
		return
	}
	signalHandler := &onExitSignal{
		exitFunc: f,
	}
	RegisterSignal(signalHandler)
}
