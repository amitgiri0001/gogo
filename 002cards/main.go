package main

func main() {
	cards := newDeck()

	hand, _ := deal(cards, 5)

	// Wite to the file
	hand.saveToFile("handFile")

	// Red from the file
	handFromFile := readFromFile("handFile")
	handFromFile.print()

	// Read a non existing file.
	handFromFile = readFromFile("handFile_1")
	handFromFile.print()
}
