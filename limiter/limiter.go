package limiter

import "log"

type Limiter interface {
	AddTask(task Task) (func(), error)
}

type LimiterDemo struct {
	ch chan struct{}
}

func NewLimiterDemo(taskNum int) Limiter {
	return &LimiterDemo{
		ch: make(chan struct{}, taskNum),
	}
}

func (l LimiterDemo) AddTask(task Task) (func(), error) {
	log.Printf("addTask pre: %+v", task)
	l.ch <- struct{}{}
	log.Printf("addTask post: %+v", task)

	return func() {
		log.Printf("delTask pre: %+v", task)
		<-l.ch
		log.Printf("delTask post: %+v", task)
	}, nil
}
