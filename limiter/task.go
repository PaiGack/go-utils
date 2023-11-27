package limiter

import (
	"fmt"
	"math/rand"
)

type Task interface {
	String() string
}

type TaskDemo struct {
	Uuid string
}

func NewTaskDemo() Task {
	return &TaskDemo{Uuid: fmt.Sprintf("%d-%d", rand.Int31n(10), rand.Int31n(10))}
}

func (t TaskDemo) String() string {
	return fmt.Sprintf("task: %s", t.Uuid)
}
