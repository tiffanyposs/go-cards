package main

import "fmt"

// Create a new type of deck
// which is a slice of strings

// create a new type called deck that
// extends the functionality of a slice of strings
type deck []string

// any variable with the type deck, gets a method named print
// (d deck) syntax is called a receiver
// a receiver sets up a method on variable we create
// the first variable "d" represents the instance of the type (kind of similar to "this" in JS)
// the second variable is the type you are extending
// by convention, the first variable in a receiver is named with 1 or 2 letters that match
// the first letters of the type it's extended, this is why we named this "d" instead of "cards"
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
