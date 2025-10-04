package advanced

import "fmt"

/*
* send only channel: chan<- type
* receive only channel: <-chan type
* you can't pass a receive only channel as an argument to a send only function parameter or vice versa
* bidirectional can be passed to both send only and receive only function parameter as an argument
 */
func directions() {
	ch := make(chan int)

	producer(ch)
	for v := range ch { // you can't range over send only channels
		fmt.Println("Received:", v)
	}

	consumer(ch)
}

// Send only channel
func producer(ch chan<- int) {
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

// Receive only channnel
func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Println("Received:", v)
	}
}
