package main

import (
	"os"
	"testing"
)

// newDeck
// What do I care about to test
// 52 cards, check the length of the deck
// First card is ""
// Last card is ""
// t is the test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades to be first, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs ot be first, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	f := "_decktesting" // set name of test file to create

	os.Remove(f) // make sure everything is cleaned up

	d := newDeck()  // create new deck
	d.saveToFile(f) // save it to file

	loadedDeck := newDeckFromFile(f) // load the file

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove(f)
}
