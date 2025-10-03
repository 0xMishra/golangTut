package main

import (
	"fmt"
	"time"
)

/* goroutines are just functions that leave the main thread and run in the background
and come back to join the main thread once the function is finished/ready to return any value
*/

// ** goroutines do not stop the program flow and are non blocking **

func goroutines() {
	var err error
	fmt.Println("before sayHello() ran")
	go sayHello()                       // this goes to the background off the main thread
	fmt.Println("after sayHello() ran") // go runtime does not wait for sayHello to finish executing

	go func() {
		err = doThisWork()
	}()
	// err =  go doWork() -- incorrect syntax ( cannot do that)

	go printNumbers()
	go printLetters()

	time.Sleep(2 * time.Second)

	// we moved it after the sleep statement so that the doWork goroutine can finish
	// otherwise there will be no error returned and this if statement will not work as expected
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("work completed successfully")
	}
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from goroutine")
}

func printNumbers() {
	for i := range 5 {
		fmt.Println(i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println(string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doThisWork() error {
	// simulate some work
	time.Sleep(1 * time.Second)
	return fmt.Errorf("an error occured in doWork function")
}
