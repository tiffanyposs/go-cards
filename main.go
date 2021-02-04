package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

// you must declare what type this function with return
func newCard() string {
	return "Five of Diamonds"
}
