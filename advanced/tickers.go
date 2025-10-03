package main

import (
	"fmt"
	"time"
)

func tickerAndTimer() {
	ticker := time.NewTicker(time.Second)
	stop := time.After(time.Duration(5) * time.Second)

	defer ticker.Stop()

	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("Tick at:", tick)
		case <-stop:
			fmt.Println("stopping ticker...")
			return
		}
	}
}

func periodicTicker() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			periodicTask()
		}
	}
}

func periodicTask() {
	fmt.Println("performing periodic task at:", time.Now())
}

func newTicker() {
	// tickers never expire, we always have to manually stop a ticker
	ticker := time.NewTicker(time.Second)

	defer ticker.Stop()

	// for tick := range ticker.C {
	// 	fmt.Println("Tick at:", tick)
	// }

	i := 1
	for range 5 {
		i *= 2
		fmt.Println(i)
	}
}
