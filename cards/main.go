package main

import "fmt"

// var card string = "Ace of Spades"

func main() {
	// var card string
	// card := "Ace of Spades"
	// card = "Five of Diamonds"

	// card := newCard()

	// fmt.Println(card)

	// cards := []string{"Ace of Spaces", newCard()}
	// cards = append(cards, "Six of Diamonds")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	cards := newDeck()
	fmt.Println(cards.toString())

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	cards.saveToFile("deck")

	deck := newDeckFromFile("deck")

	deck.shuffle()

	deck.print()

}
