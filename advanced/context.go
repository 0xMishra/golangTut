package advanced

import (
	"context"
	"fmt"
	"log"
	"time"
)

func contextTut() {
	rootCtx := context.Background()
	ctx, cancel := context.WithTimeout(
		rootCtx,
		time.Duration(1)*time.Second,
	) // timer of the context starts here, you have the specified time duration to use this context
	defer cancel()

	/* manual cancellation of context without a timer
	ctx, cancel = context.WithCancel(ctx)

	go func() {
		time.Sleep(time.Duration(2) * time.Second) - simulating a heavy task, consider this a time consuming task
		cancel() - manually cancelling only after the task is finished
	}()
	*/

	ctx = context.WithValue(ctx, "reqId", "1234")
	ctx = context.WithValue(ctx, "IP", "123.43.255.255")
	ctx = context.WithValue(ctx, "name", "john")
	ctx = context.WithValue(ctx, "age", "31")

	go doWork(ctx)
	time.Sleep(time.Duration(2) * time.Second)

	// even after 3 seconds the context will still stay
	// cancel func just sends the cancel signal, it does not destroy the context
	reqID := ctx.Value("reqId")
	if reqID != nil {
		fmt.Println(reqID)
	} else {
		fmt.Println("No request Id found")
	}

	logWithContext(ctx, "This is a test log message")
}

func logWithContext(ctx context.Context, msg string) {
	reqID := ctx.Value("reqId")
	log.Printf("Request ID: %v - %v", reqID, msg)
}

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled", ctx.Err())
			return
		default:
			fmt.Println("Working...")
		}
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
}

func timeoutContext() {
	ctx := context.TODO()

	res := checkEvenOdd(ctx, 5)
	fmt.Println("Result with context.TODO:", res)

	// context cancellation doesn't mean this context will be destroyed
	// but it means after the deadline it will send a cancellation signal
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1)*time.Second)
	defer cancel()

	res = checkEvenOdd(ctx, 10)
	fmt.Println("Result from timeout context:", res)

	time.Sleep(time.Duration(2) * time.Second)

	res = checkEvenOdd(ctx, 15)
	fmt.Println("Result after timeout:", res)
}

func checkEvenOdd(ctx context.Context, num int) string {
	select {
	case <-ctx.Done():
		return "Operation cancelled"
	default:
		if num%2 == 0 {
			return fmt.Sprintf("%d is even", num)
		} else {
			return fmt.Sprintf("%d is odd", num)
		}
	}
}

func difference() {
	/*
		context.TODO can do anything that context.Background can
		but in go it's a convention to use context.Background for storing
		actual data and context.TODO is only for placeholder or when you are unsure
		which context to use
	*/
	todoContext := context.TODO() // TODO context just used as a placeholder, doesn't hold any data

	type key string

	var key1 key = "name"
	ctx := context.WithValue(todoContext, key1, "john")
	fmt.Println(ctx)
	fmt.Println(ctx.Value(key("name")))

	var key2 key = "city"
	contextBg := context.Background()
	ctxBg := context.WithValue(contextBg, key2, "New york")
	fmt.Println(ctxBg)
	fmt.Println(ctxBg.Value(key("city")))
}
