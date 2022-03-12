package main

func main() {

	// var wale = [4]int {1, 2, 4, 5}

	strings := []string {"Hello", "Adewale Olaoye", "Playing with go"}

	strings = append(strings, "Golang Developer")

	for i, string := range strings {
		println(i, string)
	}

	// fmt.Println(wale, "Wale")


	// println(strings)
	// strings := "Hello, what's new here?"

	// for i, string := range strings {
	// 	fmt.Println(i, string)
	// }

}