package main

func main() {
	cards := newDeck()
	// fmt.Println(cards.toString())
	cards.saveToFile("my_cards")
}

// $ go run main.go deck.go
