// Author: coolliu
// Date: 2021/7/22
// 启动一组工作协程

package work

import (
	"context"
	"sync"
)

type Task func(ctx context.Context)

type Workers struct {
	group  *sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

func NewWorkers() *Workers {
	ctx, cancel := context.WithCancel(context.Background())
	return &Workers{
		group:  &sync.WaitGroup{},
		ctx:    ctx,
		cancel: cancel,
	}
}

func (w *Workers) AddWorker(task Task) {
	if task == nil {
		return
	}
	w.group.Add(1)
	go func() {
		task(w.ctx)
		w.group.Done()
	}()
}

// stop by signal
func (w *Workers) Stop() {
	if w.cancel != nil {
		w.cancel()
	}
}

func (w *Workers) Wait() {
	w.group.Wait()
}

func IsDone(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
