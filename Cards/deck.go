package main

import (
	"fmt"
	"strings"
)

//create a new type of deck
//which is a slice of string

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuit := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Fives", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuit {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
