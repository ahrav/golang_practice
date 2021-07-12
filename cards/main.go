package main

func main() {
	cards := newDeck()

	cards.saveToFile("myStuff")
	newCards := newDeckFromFile("myStuff")
	newCards.print()
	newCards.shuffle()
	newCards.print()
}
