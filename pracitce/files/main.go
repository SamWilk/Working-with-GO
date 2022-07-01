package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello From Go!"
	file, err := os.Create("./fromString.txt")
	CheckError(err)
	length, err := io.WriteString(file, content)
	CheckError(err)
	fmt.Printf("Wrote a file with %v characters\n", length)
	defer file.Close()
	defer ReadFile("./fromString.txt")
}

func ReadFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	CheckError(err)
	fmt.Println("Text from file: ", string(data))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
