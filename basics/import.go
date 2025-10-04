package basics

import (
	"fmt"
	server "net/http" // named imports
)

func imports() {
	fmt.Println("Hello Go Std. lib")

	res, err := server.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return

	}

	defer res.Body.Close()
	fmt.Println("HTTP Status: ", res.Status)
}
