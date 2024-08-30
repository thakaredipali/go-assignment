package main

import (
	"fmt"
)

func main() {
	array := []string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}
	var index1, index2 int
	fmt.Print("Enter range: ")
	fmt.Scanln(&index1, &index2)
	if index1 >= len(array) || index2 >= len(array) {
		fmt.Println("Incorrect Indexes")
		return
	}
	slice1 := array[:index1+1]
	slice2 := array[index1 : index2+1]
	slice3 := array[index2:]
	fmt.Printf("%v\n%v\n%v\n", slice1, slice2, slice3)

}
