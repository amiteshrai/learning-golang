package main

import "fmt"

func main() {
	var message string = "Running Cards Project ..."
	fmt.Println(message)
	card := newCard()
	fmt.Println("Got new card : " + card)
}

func newCard() string {
	return "Five Of Diamonds"
}
