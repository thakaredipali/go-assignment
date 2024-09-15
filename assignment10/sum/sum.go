package main

import "fmt"

func SumOfSquares(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num * num
	}
	return sum
}

func main() {
	input := []int{5, 9, 3}
	sum := SumOfSquares(input)
	fmt.Println("Output: ", sum)
}
