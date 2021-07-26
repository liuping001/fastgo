// Author: coolliu
// Date: 2021/7/22

package work

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Fun(ctx context.Context) {
	//for {
	//	if IsDone(ctx) {
	//		break
	//	}
	//	// 实际的工作
	//	time.Sleep(time.Second)
	//	fmt.Println("run")
	//}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			// 实际的工作
			time.Sleep(time.Second)
			fmt.Println("run")
		}
	}
}

func TestWork(t *testing.T) {
	work := NewWorkers()

	work.AddWorker(Fun)
	work.AddWorker(Fun)
	work.AddWorker(Fun)
	work.AddWorker(Fun)

	time.Sleep(2 * time.Second)
	work.Stop()
	work.Wait()
	fmt.Println("exit ok")
}
