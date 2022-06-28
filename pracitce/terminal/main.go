package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Sam's GO Terminal")

	//Maybe create a channel to with go streams with!

	for true {
		fmt.Print("$ ")
		//Getting User Input
		reader := bufio.NewReader(os.Stdin)
		stringInput, error := reader.ReadString('\n')
		if error != nil {

		}
		stringInput = strings.TrimSpace(stringInput)
		if strings.Contains(stringInput, "hello") == true {
			fmt.Println("Hello there")
		}
		if strings.Contains(stringInput, "goodbye") == true {
			fmt.Println("Goodbye for now")
			break
		}
	}
}
