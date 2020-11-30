package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println(target)
	fmt.Println("I've choosen a number between 1 and 100")
	fmt.Println("Can you guess it?")
	reader := bufio.NewReader(os.Stdin)
	success := false

	for guesses := 0; guesses <= 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guessses remaining!")
		fmt.Print("Make a guess:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("your guess is too low")
		} else if guess > target {
			fmt.Println("your guess is too high")
		} else {
			fmt.Println("You got it")
			success = true
			break
		}

	}

	if !success {
		fmt.Println("You lost the game")
	}
}
