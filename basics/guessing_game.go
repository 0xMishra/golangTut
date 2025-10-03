package main

import (
	"fmt"
	"math/rand"
	"time"
)

func guessingGame() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// generating a random number between 1 and 100
	target := random.Intn(100) + 1

	fmt.Println("Welcome player!")
	fmt.Println("Guess the secret number between 1 and 100")

	var guess int

	for {
		fmt.Println("Enter your guess:")
		fmt.Scanln(&guess)
		if guess == target {
			fmt.Printf("You guessed it right, the secret was %d\n", guess)
			break
		} else if guess < target {
			fmt.Println("Your guess is too low, try guessing a higher number")
		} else {
			fmt.Println("Your guess is too high, try guessing a lower number")
		}
	}
}
