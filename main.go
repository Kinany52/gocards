// Package main is the entry point for the application.
package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()
}
