package basics

import "fmt"

func rangeTut() {
	message := "hello world"

	for i, v := range message {
		fmt.Println(i, v)
		fmt.Printf("index: %d, rune: %c\n", i, v)
	}
}
