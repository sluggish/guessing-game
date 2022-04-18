package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min, max := 1, 100
	num := rand.Intn(max-min) + min
	fmt.Println("Guess a number between 1-100: ")

	for {
		var guess int
		fmt.Scanln(&guess)
		if guess == num {
			fmt.Println("You've got it! The correct number you guess is", num)
			fmt.Println("Type `exit` to exit the script")
			var exit string
			fmt.Scanln(&exit)
			os.Exit(0)
		} else if guess < num {
			fmt.Println("Higher!")
		} else if guess > num {
			fmt.Println("Lower!")
		}
	}
}