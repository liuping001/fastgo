// Author: coolliu
// Date: 2021/7/26

package work

import "context"

type ProducerConsumer struct {
	producer *Workers
	consumer *Workers
	queue    chan interface{}
}

func NewProducerConsumer(qSize int) *ProducerConsumer {
	return &ProducerConsumer{
		producer: NewWorkers(),
		consumer: NewWorkers(),
		queue:    make(chan interface{}, qSize),
	}
}

type PTask func(ctx context.Context, p chan<- interface{})
type CTask func(c <-chan interface{})

func (pc *ProducerConsumer) AddProducer(task PTask) {
	if task == nil {
		return
	}
	pc.producer.group.Add(1)
	go func() {
		task(pc.producer.ctx, pc.queue)
		pc.producer.group.Done()
	}()
}

func (pc *ProducerConsumer) AddConsumer(task CTask) {
	if task == nil {
		return
	}
	pc.consumer.group.Add(1)
	go func() {
		task(pc.queue)
		pc.consumer.group.Done()
	}()
}

func (pc *ProducerConsumer) Stop() {
	pc.producer.Stop()
	pc.producer.Wait()
	close(pc.queue)
}

func (pc *ProducerConsumer) Wait() {
	pc.consumer.Wait()
}
