package main

func main() {
	cards := []string {newWale(), practicingSlice()}

	cards = append(cards, "The appended value")

	for i, card := range cards{
		println(i, card)
	}

	for i := 1; i <= 5; i++ {
		println(i, "The value")
	};


	var n, sum = 10, 0;

	for i := 1; i < n; i++ {
		sum += i
		// sum = sum + i
	}

	println(sum, "TOTAL SUM")

	numbers := [5] int {1, 212, 32222, 40, 5}

	for item := range numbers {
		println(numbers[item])
	}
}

func newWale() string {
	return "Wale is awesome"
}

func practicingSlice() string {
	return "Practicing slice"
}