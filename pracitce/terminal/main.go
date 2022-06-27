package main

import (
	"bufio"
	"fmt"
	"os"
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
		switch stringInput {
		case "Hello":

		}
	}
}
