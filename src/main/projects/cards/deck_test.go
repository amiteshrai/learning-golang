package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGetNewDeck(t *testing.T) {
	newDeck := getNewDeck()

	if len(newDeck) != 20 {
		t.Error("Expected a deck with 20 cards, but got", len(newDeck))
	}

	// Asserting elements in a deck
	if newDeck[0] != "Ace Of Spade" {
		t.Errorf("Expected first card of 'Ace Of Spade', but got %v", newDeck[0])
	}
}

/**
* Test Save to disk and load the decks back
 */
func TestSaveToFile(t *testing.T) {
	os.Remove("decktest")

	newDeck := getNewDeck()
	newDeck.saveToFile("decktest")

	loadedDeck := getDeckFromFile("decktest")
	for i, card := range loadedDeck {
		fmt.Printf("Index: %d, Card: %v", i, card)
		fmt.Println()
	}

	if len(loadedDeck) != 20 {
		t.Error("Expected a deck with 20 cards, but got", len(loadedDeck))
	}

	os.Remove("decktest")

}
