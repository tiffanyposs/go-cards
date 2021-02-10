package main

import "fmt"

// Create a new type of deck
// which is a slice of strings

// create a new type called "deck" that
// extends the functionality of a slice of strings
type deck []string

// create a new function that returns a type "deck", which we created above
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// if there is a variable you don't need to use, replace it with an underscore
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// d deck - d is same as this, this instance of deck
// handSize - argument passed in with type int
// (deck, deck) - returns two things with type deck
// return <x>, <y> - return two things
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// any variable with the type deck, gets a method named print
// (d deck) syntax is called a receiver
// a receiver sets up a method on variable we create
// the first variable "d" represents the instance of the type (kind of similar to "this" in JS)
// the second variable is the type you are extending
// by convention, the first variable in a receiver is named with 1 or 2 letters that match
// the first letters of the type it's extended, this is why we named this "d" instead of "cards"
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
