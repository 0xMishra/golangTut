package advanced

import "fmt"

// closing channels closes the sending of data
// we can still receive data (if there's any) present only in case of buffered channels
func closingChannels() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer1(ch1)
	go filter(ch1, ch2)

	for v := range ch2 {
		fmt.Println("Received:", v)
	}
}

func producer1(ch chan<- int) {
	for i := range 5 {
		ch <- i
	}
	close(ch)
}

func filter(in <-chan int, out chan<- int) {
	for v := range in {
		if v%2 == 0 {
			out <- v
		}
	}
	close(out)
}

// loop over a closed channel
func mainC() {
	ch := make(chan int)

	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Println("Received:", val)
	}
}

func mainB() {
	ch := make(chan int)
	close(ch)

	val, ok := <-ch

	if !ok {
		fmt.Println("Channel is closed")
		return
	}

	fmt.Println("Received:", val)
}

// simple closing channel example
func mainA() {
	ch := make(chan int)

	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
