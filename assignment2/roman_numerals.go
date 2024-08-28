package main

import (
	"fmt"
)

func returnRomanValue(input string) int {

	romanValues := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var result int
	var prevValue int

	for i := 0; i < len(input); i++ {
		currentValue := romanValues[string(input[i])]
		if prevValue < currentValue && prevValue > 0 {
			result -= prevValue
			result += currentValue - prevValue
		} else {
			result += currentValue
		}

		prevValue = currentValue
	}
	return result

}

func main() {
	var input string
	fmt.Println("Enter Roman numeral")
	fmt.Scanln(&input)
	fmt.Println("Roman value", returnRomanValue(input))
}
