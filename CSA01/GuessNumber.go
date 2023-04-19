package main

import (
	"fmt"
	"math/rand"
)

func guess() {
	maxNum := 100
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
	var guess int
	for {
		_, err := fmt.Scan(&guess)
		if err != nil {
			return
		}
		if guess > secretNumber {
			fmt.Println("your guess is big")
		} else if guess < secretNumber {
			fmt.Println("your guess is small")
		} else {
			fmt.Println("your guess is right")
			break
		}
	}
}
