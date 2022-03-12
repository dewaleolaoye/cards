package main

import "fmt"

// Create a new type of deck
// which is a slice of strings

type deck []string

func (de deck) print() {
	for i, card:= range de {
		fmt.Println(i, card)
	}
}