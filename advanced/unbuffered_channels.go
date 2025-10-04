package advanced

import (
	"fmt"
	"time"
)

/* by default a channel is unbuffered, a /buffered channel means storage

* buffered channels allow asynchronous communication and will only block if the buffer is full
* unbuffered channels need to pass the value as soon they received it, it cannnot hold those values
* unbuffered channels block on send if there's no corresponding receiver, it allows all goroutine to finish
	even if those gorouitnes have nothing to do with that, then if the receiver is not present it caused error
* receiver waits for all the goroutines to finish

*/

func unbuffChannel() {
	ch := make(chan int)
	go func() {
		ch <- 1
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Println("2 second goroutine")
	}()

	go func() {
		// ch <- 1
		time.Sleep(time.Duration(3) * time.Second)
		fmt.Println("3 second goroutine")
	}()

	receiver := <-ch // blocks the execution flow until it receives a value
	fmt.Println("end of program", receiver)
}
