package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	ID   int
	Task string
}

func (w *Worker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("workerID %d started %s...\n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("workerID %d finished %s.\n", w.ID, w.Task)
}

func main() {
	tasks := []string{"digging", "laying bricks", "painting"}
	var wg sync.WaitGroup

	for i, task := range tasks {
		worker := Worker{ID: i + 1, Task: task}
		wg.Add(1)
		go worker.PerformTask(&wg)
	}

	// wait for all workers to finish
	wg.Wait()
	fmt.Println("construction is finished.")
}

func waitGrpCh() {
	var wg sync.WaitGroup

	numWorkers := 3
	numJobs := 5
	results := make(chan int, numJobs)
	tasks := make(chan int, numJobs)

	wg.Add(numWorkers)

	for i := range numWorkers {
		go workerCh(i+1, tasks, results, &wg)
	}

	for i := range numJobs {
		tasks <- i + 1
	}
	close(tasks)

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println("Result:", res)
	}
}

func workerCh(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d is starting...\n", id)
	time.Sleep(time.Second) // simulating some work
	for task := range tasks {
		results <- task * 2
	}
	fmt.Printf("worker %d finished.\n", id)
}

func workerB(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// wg.Add(1) -- WRONG PRACTICE, never use Add func inside a goroutine
	fmt.Printf("worker %d starting...\n", id)
	time.Sleep(time.Second) // simulating time spent on processing
	fmt.Printf("worker %d finished.\n", id)
}

func waitGrp() {
	var wg sync.WaitGroup
	numWorkers := 3
	wg.Add(numWorkers)

	// launch the workers
	for i := range numWorkers {
		go workerB(i, &wg)
	}

	wg.Wait() // blocking untill all goroutines are done
	fmt.Println("all workers are finished.")
}
