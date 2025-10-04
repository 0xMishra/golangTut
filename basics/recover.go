package basics

import "fmt"

func recoverTut() {
	recoverProcess()
	fmt.Println("exited from process")
}

func recoverProcess() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
		}
	}()

	fmt.Println("process started")
	panic("Something went wrong")
}
