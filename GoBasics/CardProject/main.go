package main

import "fmt"

func main() {
	cards := createDeckFromFile("Cards1")
	// cards.print()
	cards.shuffle()
	fmt.Println("<--- doing shuffle 💳")
	cards.print()

}
