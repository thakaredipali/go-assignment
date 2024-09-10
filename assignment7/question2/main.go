package main

import (
	"fmt"
	"strings"
)

var goroutineCount int

func main() {
	input := "test123"
	reversedStr := make(chan string)
	goroutineCount++
	go reverseString(input, reversedStr)
	msg := <-reversedStr
	fmt.Println(msg, goroutineCount)
}

func reverseString(s string, ch chan string) {
	res := ""
	var reversedStr = strings.Split(s, "")
	for i := len(reversedStr) - 1; i >= 0; i-- {
		res = res + reversedStr[i]
	}
	goroutineCount++
	ch <- res
	close(ch)
}
