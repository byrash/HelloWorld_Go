package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	// card := newCard()
	// cards := deck{"Ace of Diamonds", newCard()}
	cards := newDeck()
	cards.shuffle()
	// hand, remainingCards := deal(cards, 5)
	// cards = append(cards, "Six of Spades")
	// fmt.Println(cards)
	// for index, card := range cards {
	// 	fmt.Println(index, card)
	// }
	// cards.print()
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards1")
	cards.print()
}

// func newCard() string {
// 	return "Five Of Diamonds"
// }
