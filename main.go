package main

func main() {
	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// fmt.Println(cards)
}

// $ go run main.go deck.go
