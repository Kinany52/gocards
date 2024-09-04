// Package main is the entry point for the application.
package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
}
