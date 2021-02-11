package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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
	for i, card := range d {
		fmt.Println(i, card)
	}
	fmt.Println("-----------")
}

// takes a receiver of d deck, function named toString, returns a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// takes a receiver of d deck, function named saveToFile, takes a param of filename, returns an error type
// cards.saveToFile("my_cards")
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // 0666 - anyone can use this file
}

// reads a file on os
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	// if there is an error, print error and exit program completely
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1) // exit program by passing any number besides 0
	}

	s := strings.Split(string(bs), ",") // turn byte slice into a string and split it into a slice of strings
	return deck(s)                      // convert it into a deck type
}

func (d deck) shuffle() {
	// Go's rand.Intn is a pseudo random generator, so it's the same every time by default

	// to create a true random number
	// create a random number by using the current time
	source := rand.NewSource(time.Now().UnixNano())
	// create new random number generator
	r := rand.New(source) // type Rand

	for i := range d {
		// rand.Intn(len(d) - 1) - not truely random, has same seed each time
		// create random number with new Rand type we created
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
