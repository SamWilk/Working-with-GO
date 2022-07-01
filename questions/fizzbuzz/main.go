package main

import "fmt"

func main() {

	fizzBuzz := make(map[int]string)
	number := [2]int{3, 5}
	var count int = 1
	for count <= 40 {
		hitCount := 0
		for _, value := range number {
			switch count % value {
			case 0:
				if value == 3 {
					fizzBuzz[count] = "fizz"
					hitCount++
				} else {
					if hitCount >= 1 {
						fizzBuzz[count] = "fizz buzz"
					} else {
						fizzBuzz[count] = "buzz"
					}
				}
			default:
				break
			}
		}
		fmt.Println(count, ":", fizzBuzz[count])
		count++
	}

}
