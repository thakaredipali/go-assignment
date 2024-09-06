package main

import (
	"errors"
	"fmt"
)

func accessSlice(slice []int, index int) error {

	if index > len(slice)-1 {
		return errors.New("length of the slice should be more than index")
	}
	fmt.Printf("item %d, value %d\n", index, slice[index])
	return nil

}

func main() {
	slice := []int{1, 2, 4, 6, 9, 23, 56}

	var index int
	fmt.Println("Enter index")
	fmt.Scanln(&index)
	err := accessSlice(slice, index)
	if err != nil {
		fmt.Println(err)
	}
}
