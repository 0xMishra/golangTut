package basics

import (
	"fmt"
	"os"
)

func exitTut() {
	defer fmt.Println(
		"deferred statement",
	) // os.Exit exits without any cleanups so this will never run
	fmt.Println("start of main function")

	os.Exit(0) // avoids defer statement

	// this will never run
	fmt.Println("end of main function")
}
