package main

func main() {
	cards := newDeck()
	cards.print()
	cards.shuffle()
	println("Shuffled.")
	cards.print()
	//cards := newDeckFromFile("my_cards.txt")

	//cards := newDeckFromFile("rr.txt")
	//cards := newDeck()
	//cards.saveToFile("my_cards.txt")
	//fmt.Println(cards.toString())
	//

	//hand, remainingDeck := deal(cards, 5)
	//
	//hand.print()
	//remainingDeck.print()
	//greetings := "hi there"
	//fmt.Println([]byte(greetings))

}
