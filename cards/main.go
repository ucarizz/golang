package main

import "fmt"

func main() {
	card := []string{newCard(), newCard()}
	card = append(card, "yeni card")
	//var = We're about to create a new variable
	//card = The name of the variable will be 'greeting'
	//string = Only a "string" will ever be assigned to this variable
	//"Ace of Spades" = Assign the value "Ace of Spades" to this variable
	fmt.Println(card)

}

func newCard() string {
	return "TestCard"
}
