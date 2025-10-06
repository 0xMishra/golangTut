package advanced

import (
	"fmt"
	"time"
)

type StatefulWorker struct {
	count int
	ch    chan int
}

func (w *StatefulWorker) start() {
	go func() {
		for {
			select {
			case value := <-w.ch:
				w.count += value
				fmt.Println("current count:", w.count)
			}
		}
	}()
}

func (w *StatefulWorker) Send(value int) {
	w.ch <- value
}

func statefulGoroutine() {
	stWorker := &StatefulWorker{
		ch: make(chan int),
	}

	stWorker.start()

	for i := range 5 {
		stWorker.Send(i)
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
}
