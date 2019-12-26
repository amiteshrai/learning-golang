/**
 * Package main
 */

package main

import "fmt"

// Create on new type of 'deck' which is a slice of string
type deck []string

// Create a function to print elements of deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, " : ", card)
	}
}

// Create a deck and return
func getDeck() deck {
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
