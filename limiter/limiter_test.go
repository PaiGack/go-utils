package limiter

import (
	"testing"
	"time"
)

func TestLimiterDemo(t *testing.T) {
	limiter := NewLimiterDemo(2)
	t1, _ := limiter.AddTask(NewTaskDemo())
	t2, _ := limiter.AddTask(NewTaskDemo())

	go func() {
		time.Sleep(time.Second * 10)
		t1()
	}()

	t3, _ := limiter.AddTask(NewTaskDemo())
	t3()

	t4, _ := limiter.AddTask(NewTaskDemo())

	t2()
	t4()
}

func TestLimiterDemoFail(t *testing.T) {
	limiter := NewLimiterDemo(2)
	t1, _ := limiter.AddTask(NewTaskDemo())
	t2, _ := limiter.AddTask(NewTaskDemo())
	t3, _ := limiter.AddTask(NewTaskDemo())

	t3()
	t1()
	t2()
}
