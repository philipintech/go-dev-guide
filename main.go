package main

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// cards.print()
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// fmt.Println("Remaining Hand is....")
	// remainingCards.print()
	// cards := newDeck()
	// cards.saveToFile("My_cards")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// continue at video #31
