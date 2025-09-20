package main

import "fmt"

func main() {
	// var card string = "Something"
	card := giveString()
	fmt.Println(card)

}

func giveString() string {
	return " Some String is [passed]"
}
