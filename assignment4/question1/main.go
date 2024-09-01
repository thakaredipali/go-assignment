package main

import (
	"fmt"
)

type Details struct {
	name  string
	email string
}

func AcceptAnything(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("This is a value of type Integer, %v\n", v)
	case string:
		fmt.Printf("This is a value of type String, %v\n", v)
	case bool:
		fmt.Printf("This is a value of type Boolean, %v\n", v)
	case Details:
		fmt.Printf("This is a value of type Custom, Name: %v, Email: %v\n", v.name, v.email)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	var number int
	fmt.Println("Enter a number between 1 and 4:")
	fmt.Scanln(&number)

	switch number {
	case 1:
		var number int
		fmt.Println("Send integer type of object")
		fmt.Scanln(&number)
		AcceptAnything(number)
	case 2:
		var word string
		fmt.Println("Send string type of object")
		fmt.Scanln(&word)
		AcceptAnything(word)
	case 3:
		var value bool
		fmt.Println("Send boolean type of object")
		fmt.Scanln(&value)
		AcceptAnything(value)
	case 4:
		var cutomeObject Details
		fmt.Println("Send custom type of object as name and email")
		fmt.Scanln(&cutomeObject.name, &cutomeObject.email)
		AcceptAnything(Details(cutomeObject))
	default:
		fmt.Println("Invalid number")
	}
}
