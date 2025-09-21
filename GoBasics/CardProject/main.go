package main

import (
	"fmt"
	"slices"
)

func main() {
	// var card string = "Something"
	cards := deck{"A", "b", "c"}
	card := giveString()
	fmt.Println(card)
	fmt.Println(cards)
	cards = append(cards, "d")
	fmt.Println(cards)
	fmt.Println("<--------------------->")
	for i, car := range cards {
		fmt.Println("i", i, " card : ", car)
	}

	fmt.Printf("cards: %v\n", cards)

	slices.Reverse(cards)

	replaced := slices.Replace(cards, 1, 2)
	fmt.Println("replaced", replaced)

	fmt.Printf("cards: %v\n", cards)

	// Arrays

	arr := [3]int{1, 2, 3}

	fmt.Println(arr)
	fmt.Printf("arr: %v\n", arr)
	fmt.Println(len(arr))

	cards.print()

}

func giveString() string {
	return " Some String is [passed]"
}
