// Package main is the entry point for the application.
package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	fmt.Println("________________")
	remainingCards.print()
}
