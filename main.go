package main

func main() {
	cards := deck {"Ace of Diamonds", newCard()}
	cards = append(cards, "The appended value")

	// for i, card := range cards{
	// 	println(i, card)
	// }

	cards.print();
}



func newCard() string {
	return "Five of Diamonds"
}