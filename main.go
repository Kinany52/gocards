// Package main is the entry point for the application.
package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	//outputs the index and the cards as per the loop
	for i, card := range cards {
		fmt.Println(i, card)
	}

	//will not output the actual values of the elements in the slice. Instead, it will iterate over the indices and outputs the indices of the elements
	for card := range cards {
		fmt.Println(card)
	}

	//to print the actual values (i.e., the card names), you need to use both the index and value.
	//the _ (underscore) is used to ignore the index since it’s not needed.
	//'card' receives the value of each element in the cards slice.
	for _, card := range cards {
		fmt.Println(card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
