package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardsSuits := []string{"Sapdes", "Hearts"}
	cardsValues := []string{"Ace", "Two", "Three"}

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
	for i, card := range d {
		fmt.Println(i, card)
	}
	fmt.Println()
}

/**
 *
 */
func deal(d deck, handsNumber int) (deck, deck) {
	return d[:handsNumber], d[handsNumber:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filepath string) deck {
	bytes, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bytes), ",")
	return deck(s)
}
