/**
 * Package main
 */

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create on new type of 'deck' which is a slice of string
type deck []string

// Create a function to print elements of deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, " : ", card)
	}
}

// Create a deck and return it
func getNewDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spade", "Diamond", "Heart", "Club"}
	cardValues := []string{"Ace", "2", "3", "4", "5"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" Of "+suite)
		}
	}
	return cards
}

/**
* Deal cards with given handsize and return two values
 */
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString(seperator string) string {
	deckStr := ""
	for _, card := range d {
		deckStr = deckStr + seperator + card
	}

	return deckStr
}

func (d deck) saveToFile(filename string) error {
	// deckStr := strings.Join(d, "|")
	deckStr := d.toString("|")
	deckStrByte := []byte(deckStr)
	return ioutil.WriteFile(filename, deckStrByte, 0666)
}

func getDeckFromFile(filename string) deck {
	deckStrByte, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option 1: log the error and return a call to newDeck
		// Option 2: log the error and quit
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	deckStr := strings.Split(string(deckStrByte), "|")

	return deck(deckStr)
}

// Shuffle all the elements in a slice
func (d deck) shuffleDeck() deck {
	endIndex := len(d) - 1
	currentTime := time.Now().UnixNano()
	source := rand.NewSource(currentTime)
	r := rand.New(source)

	// rand.Seed(100)
	for currentIndex := range d {
		newIndex := r.Intn(endIndex)
		d[currentIndex], d[newIndex] = d[newIndex], d[currentIndex]
	}

	return d
}
