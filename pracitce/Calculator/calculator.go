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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Value: ")
	input1, _ := reader.ReadString('\n')
	num1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Entry Not allowed")
	}

	fmt.Print("Value: ")
	input2, _ := reader.ReadString('\n')
	num2, err2 := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err2 != nil {
		fmt.Println(err2)
		panic("Entry Not allowed")
	}
	num1 = num1 + num2
	fmt.Println("Final Value:", num1)
}
