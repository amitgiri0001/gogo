package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardsSuits := []string{"Sapdes", "Hearts", "Diamonds", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "king", "queen", "jack"}

	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

/**
 * Receiver block to attach method to the deck type.
 */
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
