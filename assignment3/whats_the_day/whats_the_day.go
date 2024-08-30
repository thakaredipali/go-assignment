package main

import (
	"fmt"
)

func whatsTheDay(number int) string {
	days := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}
	var result string
	for index, day := range days {
		if number == index {
			result = day
			break
		} else {
			result = "Not a day"
		}
	}
	return result
}

func main() {
	var index int
	fmt.Println("Enter number")
	fmt.Scanln(&index)
	fmt.Println("What's The Day: ", whatsTheDay(index))

}
