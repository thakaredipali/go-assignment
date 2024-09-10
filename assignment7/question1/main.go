package main

import (
	"fmt"
	"sync"
)

func main() {
	aliceChan := make(chan string)
	bobChan := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go processAliceMessages(aliceChan, &wg)
	go processBobMessages(bobChan, &wg)

	input := "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"

	var currentMessage string
	for _, char := range input {
		switch char {
		case '$':
			aliceChan <- currentMessage
			currentMessage = ""
		case '#':
			bobChan <- currentMessage
			currentMessage = ""
		case '^':
			close(aliceChan)
			close(bobChan)
			wg.Wait()
		default:
			currentMessage += string(char)
		}
	}
}

func processAliceMessages(aliceChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range aliceChan {
		fmt.Println("alice :", msg)
	}
}

func processBobMessages(bobChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range bobChan {
		fmt.Println("bob :", msg)
	}
}
