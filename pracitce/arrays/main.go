package main

import (
	"fmt"
	"sort"
)

func main() {
	//Arrays
	var numbers = [4]int{1, 4, 5, 3}
	fmt.Println(numbers)
	//Slices
	//They can be resized and can be sorted easily
	var colors = []string{"red", "blue", "black"}
	colors = append(colors, "green")
	fmt.Println(colors)
	fmt.Println(colors[len(colors)-1:])
	var sortNumbers = []int{2, 4, 3, 1, 7, 8, 9, 65, 34, 3, 5, 67}
	sort.Ints(sortNumbers)
	fmt.Print(sortNumbers)
}
