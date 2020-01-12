/*
* Implementing guess games
**/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	const maxAttempts = 5
	const min, max = 1, 100
	userAttempts := 0
	fmt.Println("-> Guess a number between 1 and 100")

	for {
		fmt.Println("-> Attempts available : ", (maxAttempts - userAttempts))
		userAttempts++
		fmt.Print("-> Please input your guess : ")

		// Read input from console
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		// convert CRLF to LF
		input = strings.Replace(input, "\n", "", -1)

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("-> Invalid input. Please enter an integer value.")
			return
		}

		randomNumber := generateRandomInteger(min, max)

		fmt.Printf("-> Expected input : %v, Your input : %v", randomNumber, guess)
		fmt.Println()
		if randomNumber == guess {
			fmt.Println("-> Your guess is correct :)")
			break
		} else {
			fmt.Println("-> Your guess is incorrect. Please try again!")
		}

		if userAttempts >= maxAttempts {
			fmt.Println("-> You have used all the attemps available")
			fmt.Println("-> Exiting application")
			break
		}
	}

}

// Random number generator
func generateRandomInteger(min, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min) + min
}
