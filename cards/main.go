package main

func main() {
	d := newDeck()
	f := "my_deck"

	d.writeToFile(f)

	readDeckFromFile(f).print()
}
