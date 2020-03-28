package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	cards := []string{"Ace of Diamonds", newCard()}
	cards=append((cards,"six of psrades"))
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
