package advanced

import (
	"fmt"
	"time"
)

func nonBlockingOp() {
	// Non blocking receive operation
	ch := make(chan int)

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No messages available")
	}

	// Non blocking send operation
	select {
	case ch <- 1:
		fmt.Println("Sent message")
	default:
		fmt.Println("channel is not ready to receive")
	}

	// blocking operations in real time systems
	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case d := <-data:
				fmt.Println("Data received:", d)
			case <-quit:
				fmt.Println("Stopping...")
				return
			default:
				fmt.Println("Waiting for data")
				time.Sleep(time.Duration(500) * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(time.Duration(1) * time.Second)
	}

	quit <- true
}
