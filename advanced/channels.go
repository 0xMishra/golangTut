package main

import (
	"fmt"
	"time"
)

/*
	unbuffered channels don't work without goroutines present in the program

a goroutine should be there (**before the channel receives the value**) to receive values from the unbuffered channel immediately
*/
func channels() {
	// creating a channel: variable := make(chan type)
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString // blocking statement because it is continuously trying to receive data

		for _, v := range "abcde" {
			greeting <- string(v)
		}
	}()

	go func() {
		receiver := <-greeting // non blocking because receiver is part of the main goroutine, only blocking if there's no value to receive
		fmt.Println(receiver)

		for range 5 {
			receiver = <-greeting
			fmt.Println(receiver)
		}
	}()

	time.Sleep(time.Duration(50) * time.Second)

	fmt.Println("end of program")
}
