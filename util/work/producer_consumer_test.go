// Author: coolliu
// Date: 2021/7/22

package work

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func producer(ctx context.Context, queue chan<- interface{}) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// 实际的工作
			time.Sleep(time.Second)
			queue <- string("run")
		}
	}
}

func consumer(queue <-chan interface{}) {
	for {
		select {
		case msg := <-queue:
			if msg == nil {
				return
			}
			fmt.Println(msg)
		}
	}
}

func TestProducerConsumer(t *testing.T) {
	work := NewProducerConsumer(100)

	work.AddProducer(producer)
	work.AddConsumer(consumer)

	time.Sleep(2 * time.Second)
	work.Stop()
	work.Wait()
	fmt.Println("exit ok")
}
