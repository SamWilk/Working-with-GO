package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
		This is showing how to get input from a user from the command line
		the packages needed will auto import and will be removed as needed
	*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")
	/*
		The underscore will be used later on when you want to work with error handling
		Since their is no try catch type handling in GO the method will return an error object if something fails
		A _ will ignore variables if you do not want to mess with them atm
	*/
	input, _ := reader.ReadString('\n')
	fmt.Print("You Entered: ", input)

	//String Convertion
	fmt.Print("Enter a Number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	//A form or error handling
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of Number: ", aFloat)
	}
}
