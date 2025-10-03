package main

import (
	"fmt"
	"time"
)

func multipleTimers() {
	timer1 := time.NewTimer(time.Duration(1) * time.Second)
	timer2 := time.NewTimer(time.Duration(2) * time.Second)

	select {
	case <-timer1.C:
		fmt.Println("timer1 expired")
	case <-timer2.C:
		fmt.Println("timer2 expired")
	}
}

func newTimer() {
	timer := time.NewTimer(time.Duration(2) * time.Second) // non blocking timer starts

	go func() {
		<-timer.C
		fmt.Println("Delayed operation executed")
	}()

	fmt.Println("Waiting...")
	time.Sleep(time.Duration(3) * time.Second) // blocking timer starts
	fmt.Println("End of the program")
}

func timeAfter() {
	timeout := time.After(time.Duration(2) * time.Second)
	done := make(chan bool)

	go func() {
		longRunningOp()
		done <- true
	}()

	select {
	case <-timeout:
		fmt.Println("operation timed out")
	case <-done:
		fmt.Println("operation completed")
	}
}

func longRunningOp() {
	for i := range 20 {
		fmt.Println(i)
		time.Sleep(time.Second) // simulating a time consuming task
	}
}

func timer1() {
	fmt.Println("Starting app...")
	timer := time.NewTimer(time.Duration(2) * time.Second) // non blocking in nature

	fmt.Println("Waiting for timer.C")

	stopped := timer.Stop() // stoping the timer manually
	if stopped {
		fmt.Println("Timer has stopped")
	}

	timer.Reset(time.Second)
	fmt.Println("timer reset")

	<-timer.C // blocking - waiting for send channel, will expire after 2 seconds
	fmt.Println("Timer expired")
}
