package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("./file.txt")
	cards := newDeckFromFile("./file.txt")
	cards.print()
}
