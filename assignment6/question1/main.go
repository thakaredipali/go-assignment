package main

import (
	"fmt"
)

func accessSlice(slice []int, index int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("internal error: %v\n", r)
		}
	}()
	fmt.Printf("item %d, value %d\n", index, slice[index])
}

func main() {
	slice := []int{1, 2, 4, 6, 9, 23, 56}

	var index int
	fmt.Println("Enter index")
	fmt.Scanln(&index)
	accessSlice(slice, index)
	fmt.Println("Testing Panic and Recover")
}
