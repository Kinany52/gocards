// Package main is the entry point for the application.
package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
