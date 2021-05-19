package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Guess the number!")

	// generate a random number
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(100) // generates numbers between 0 and n (10)
	attempt := 1
	var boom string
	var name string
	fmt.Println("What is your name ? ")
	fmt.Scan(&name)

	for {
		fmt.Println(name, ",Please input your guess: ")
		fmt.Scan(&boom)
		guess, err := strconv.Atoi(boom)
		if err != nil {
			fmt.Println(name, ",You should enter ONLY digits!!!")

		} else if guess > secretNumber {
			fmt.Println("Too Big")
			fmt.Println("")

			attempt++
		} else if guess < secretNumber {
			fmt.Println("Too small")
			fmt.Println("")

			attempt++

		} else {
			fmt.Println(name, ",You won! 🥳 . You guessed with ", attempt, "attempts")
			break
		}
	}
}
