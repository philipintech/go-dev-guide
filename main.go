package main

import "fmt"

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// cards.print()
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	fmt.Println("Remaining Hand is....")
	remainingCards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }

// finished at video # 23
