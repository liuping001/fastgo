// Author: coolliu
// Date: 2021/5/21

package app

import (
	"fmt"
	"os"
	"syscall"
	"testing"
	"time"
)

type MyApp struct {
	DefaultOnSignal
}

func (a *MyApp) OnExit() {
	time.Sleep(1 * time.Second)
	fmt.Printf("退出程序\n")
	os.Exit(0)
}

func TestGracefulExit(t *testing.T) {
	fmt.Printf("测试优雅退出：TestGracefulExit\n")
	myApp := MyApp{DefaultOnSignal{}}
	GracefulExit(&myApp, nil)
	syscall.Kill(syscall.Getpid(), syscall.SIGUSR1)
	time.Sleep(1 * time.Second) // 需要间隔一段时间才能发下一个信号，不然会被覆盖
	syscall.Kill(syscall.Getpid(), syscall.SIGUSR2)
	time.Sleep(1 * time.Second)
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	time.Sleep(10 * time.Second)
}
