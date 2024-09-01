package main

import "fmt"

type Rectangle struct {
	length int
	width  int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func main() {
	var length, width int
	fmt.Println("Enter a lenght and width:")
	fmt.Scanln(&length, &width)
	rectangle := Rectangle{length: length, width: width}
	fmt.Println("Area of Rectangle", rectangle.Area())
	fmt.Println("Perimeter of Rectangle", rectangle.Perimeter())
}
