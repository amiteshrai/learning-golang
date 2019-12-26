package main

import "fmt"

func main() {
	var message string = "Running Cards Project ..."
	fmt.Println(message)
	card := newCard()
	fmt.Println("Got new card : " + card)

	fmt.Println("\nUsing Built In Types")

	var cards []string
	cards = append(cards, "Ace Of Spade")
	fmt.Println("Available Cards: ")
	fmt.Println(cards)
	cards = append(cards, newCard())
	fmt.Println("Available Cards: ")
	fmt.Println(cards)
	printCards(cards)

	fmt.Println("\nExtending Built In Types")
	cardsDeck := getDeck()
	fmt.Println(cardsDeck)
}

func newCard() string {
	return "Five Of Diamonds"
}

func getCards() []string {
	return []string{"Ace Of Spade"}
}

func printCards(cards []string) {
	fmt.Println("Printing Available Cards: ")
	for index, card := range cards {
		// fmt.Println(card) // Every declared variable should be used
		fmt.Println(index, card)
	}
}
