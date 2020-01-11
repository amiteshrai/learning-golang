package main

import (
	"fmt"
	"strings"
)

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
	cardsDeck := getNewDeck()
	fmt.Println(cardsDeck)

	fmt.Println("\n Selecting Subset From Slice")
	fmt.Println(cardsDeck[:5])

	fmt.Println("/nDeal Cards")
	handSize := 5
	hand, remainingDeck := deal(cardsDeck, handSize)
	fmt.Println("\nPrinting Hand Cards: ")
	hand.print()
	fmt.Println("\nPrinting Remaining Cards: ")
	remainingDeck.print()

	fmt.Println("\nType Conversion")
	fmt.Println([]byte("Hi Golang"))

	// Converting slice of string to string
	cardsDeckStr := cardsDeck.toString("|")
	fmt.Println(cardsDeckStr)
	cardsDeckStr = strings.Join(cardsDeck, "|")
	fmt.Println(cardsDeckStr)

	fmt.Println("\nSaving Deck On Disk")
	cardsDeck.saveToFile("cards.txt")

	fmt.Println("\nLoading Deck From Disk")
	deckFromFile := getDeckFromFile("cards.txt")
	fmt.Println(deckFromFile)
	// getDeckFromFile("card2.txt")

	fmt.Println("\nShuffle Deck")
	fmt.Println(deckFromFile.shuffleDeck())
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
