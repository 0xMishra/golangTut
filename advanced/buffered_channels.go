package advanced

import (
	"fmt"
	"time"
)

/*
* buffered channel only needs goroutines once they are full
* if they are not full you can also receive values from buffered channle in the main goroutine
* buffered channels block on send only if the buffer is full
* buffer channel also block on receiver only if the buffer is empty
 */

func bufChannels() {
	// variable := make(chan type, capacity)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Println("received", <-ch)
	}()
	ch <- 3 // blocks because the buffer is full
	fmt.Println("received", <-ch)
	fmt.Println("buffered channel")
	receiver()
}

func receiver() {
	ch := make(chan int, 2)
	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		ch <- 1
	}()
	fmt.Println("received:", <-ch) // blocks the thread because the buffer is empty
	fmt.Println("end of program")
}
