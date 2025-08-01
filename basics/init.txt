package main

import "fmt"

// this functions will run sequentially
func init() {
	fmt.Println("initializing the package1...")
}

func init() {
	fmt.Println("initializing the package2...")
}

func init() {
	fmt.Println("initializing the package3...")
}

func init() {
	fmt.Println("initializing the package4...")
}

func main() {
	fmt.Println("inside the main function")
}
