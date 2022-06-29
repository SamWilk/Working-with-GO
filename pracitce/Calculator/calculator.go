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
	input1 := AssignValue()
	input2 := AssignValue()

	result := GetOp(&input1, &input2)
	fmt.Println("Result:", result, "for" , input1, "and", input2)
}

func GetOp(value1 ,value2 *int) int{
Restart:
	fmt.Print("Operator: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	switch input = strings.TrimSpace(input); input {
	case "*":
		return *value1 * *value2
	case "+":
		return *value1 + *value2
	case "-":
		return *value1 - *value2
	case "/":
		return *value1 / *value2
	default:
		fmt.Println("Error,", input, "is not allowed")
		goto Restart
	}
}

func AssignValue() int {
Restart:
	fmt.Print("Value: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	temp, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println(err)
		goto Restart
	}
	return temp
}
