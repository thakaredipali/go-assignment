package main

import (
	"fmt"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	n := 1

	res := make(chan bool)
	go func() {

		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			res <- true
			return
		}
		fmt.Println(n, "is odd")
		res <- false

	}()
	<-res
	go func() {
		n++
	}()

	time.Sleep(time.Second)
}
