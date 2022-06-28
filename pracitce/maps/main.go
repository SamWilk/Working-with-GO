package main

import (
	"fmt"
	"sort"
)

func main() {
	//They are essentially Hash tables with key value pairs
	states := make(map[string]string)
	states["FL"] = "Florida"
	states["TX"] = "Texas"
	states["OR"] = "Oregon"
	states["MN"] = "Maine"
	states["CA"] = "California"
	fmt.Println(states)
	//Make will create the object and allocate memory for it while
	//new creates the object but does not allocate memory for it
	numbers := make(map[int]int)
	for i := 0; i < 100; i++ {
		numbers[i] = i * 3
	}
	//Prints out the key value pairs of the map
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}
	//Prepping the keys to be sorted from said map
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	//Sorting keys to display map in a ordered fashion
	sort.Strings(keys)
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
