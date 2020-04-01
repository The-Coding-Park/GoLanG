package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected length is 16, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card not found, got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Not found expected, got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")
	deck := newDeck()
	deck.saveToFile("_deckTesting")
	loadedDeck := newDeckFromFile("_decktesTing")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected length is 20, but got %v", len(loadedDeck))
	}
	os.Remove(("_deckTesting"))

}
