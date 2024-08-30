package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	words := takeInput()

	var result = map[string]int{}
	var newArray []string

	for i := 0; i < len(words); i++ {
		if _, exists := result[words[i]]; exists {
			fmt.Println(words[i])
			newArray = append(newArray, words[i])
		} else {
			result[words[i]] = 1
		}
	}

	fmt.Println("Highest frequency of words:", newArray)
}

func takeInput() []string {

	fmt.Println("Enter sentence")
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	wordSlice := strings.Fields(line)
	return wordSlice
}
