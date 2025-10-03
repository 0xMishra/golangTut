package main

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID   int
	numTickets int
	cost       int
}

func tickets() {
	numReq := 5
	price := 5
	ticketReq := make(chan ticketRequest, numReq)
	ticketRes := make(chan int)

	// start ticket processing
	for range 3 {
		go ticketProcessor(ticketReq, ticketRes)
	}

	// send ticket requests
	for i := range numReq {
		ticketReq <- ticketRequest{
			personID:   i + 1,
			numTickets: (i + 1) * 2,
			cost:       (i + 1) * price,
		}
	}

	close(ticketReq)

	for range numReq {
		fmt.Printf("ticket for personID %d processed succesfully\n", <-ticketRes)
	}
}

func ticketProcessor(requests <-chan ticketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf(
			"processing %d ticket(s) of personID %d with total cost %d\n",
			req.numTickets,
			req.personID,
			req.cost,
		)
		// simulating processing time of tickets
		time.Sleep(time.Second)
		results <- req.personID
	}
}

func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(time.Second) // simulating some work
		results <- task * 2
	}
}

func workerPoolPattern() {
	numWorkers := 3
	numJobs := 10
	tasks := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// create workers
	for i := range numWorkers {
		go worker(i, tasks, results)
	}

	// send tasks
	for i := range numJobs {
		tasks <- i
	}

	close(tasks)

	// collect the results
	for range numJobs {
		res := <-results
		fmt.Println("Result:", res)
	}
}
