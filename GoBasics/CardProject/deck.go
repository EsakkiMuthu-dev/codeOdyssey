package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	var cards deck
	cardSuites := []string{"Spade", "Heart", "Diamonds", "clubs"}
	cardvalues := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, cardSuite := range cardSuites {
		for _, cardvalue := range cardvalues {
			cards = append(cards, cardSuite+" of "+cardvalue)
		}

	}

	return cards
}

func deal(cards deck, handSize int) (deck, deck) {
	return cards[:handSize], cards[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToNewFile(filename string) {
	os.WriteFile(filename, []byte(d.toString()), 0666)
}

func createDeckFromFile(filename string) deck {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error : ‼️ ", err)
	}

	return deck(strings.Split(string(bytes), ","))

}

func (d deck) shuffle() {
	for index := range d {
		swapIndex := rand.Intn(len(d))
		temp := d[index]
		d[index] = d[swapIndex]
		d[swapIndex] = temp
	}
}
