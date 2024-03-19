package main

func main() {
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("s")
	cards := newDeckFromFile("my_cards")
	cards.print()
	cards.shuffle()
	cards.print()
}
