package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello from calculator")
	//Getting User Input
	input1 := assignValue()
	input2 := assignValue()

}

func assignValue() int {
	fmt.Print("Value: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	temp, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println(err)
		panic("Entry Not allowed")
	}
	return temp
}
