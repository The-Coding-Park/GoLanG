package main

//import "fmt"

func main() {
	cards := newDeck()
	//fmt.Println(cards)
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

}
