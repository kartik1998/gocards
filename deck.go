package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}
func (cards deck) print() {
	for _, card := range cards {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
