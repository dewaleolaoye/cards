package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p Person) print() {
	fmt.Printf("%s is of %d years old\n", p.name, p.age)
}

func main() {

	// strings := []string {"Hello", "Adewale Olaoye", "Playing with go"}
	// strings = append(strings, "Golang Developer")

	// for i, string := range strings {
	// 	println(i, string)
	// }

	wale := Person {
		name: "Adewale Olaoye",
		age: 45,
	}

	wale.print()

	// fmt.Println(wale)

}