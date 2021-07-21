package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("./file.txt")
	// cards := newDeckFromFile("./file.txt")
	// cards.print()
}
