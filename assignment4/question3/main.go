package main

import "fmt"

type Rectangle struct {
	length int
	width  int
}

type Square struct {
	side int
}

type Quadrilateral interface {
	Area() int
	Perimeter() int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}
func (s Square) Area() int {
	return s.side * s.side
}

func (s Square) Perimeter() int {
	return 4 * s.side
}

func Print(q Quadrilateral) {
	fmt.Printf("Area : %d\n", q.Area())
	fmt.Printf("Perimeter : %d\n", q.Perimeter())
}

func main() {
	var number int
	fmt.Println("Enter 1 for rectangle and 2 for square:")
	fmt.Scanln(&number)

	if number == 1 {
		rectangle := Rectangle{length: 30, width: 20}
		Print(rectangle)
	} else if number == 2 {
		square := Square{side: 20}
		Print(square)
	} else {
		fmt.Println("Invalid choice")
	}
}
