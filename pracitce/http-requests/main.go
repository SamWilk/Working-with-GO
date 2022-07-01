package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	//runs get request
	//Returns a pointer, not string or json
	response, err := http.Get("http://localhost:3000/api/corals")
	CheckError(err)
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	CheckError(err)
	content := string(bytes)
	fmt.Println(content)

	corals := CoralsFromJson(content)
	for _, coral := range corals {
		fmt.Println(coral.name)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CoralsFromJson(content string) []Coral {
	corals := []Coral{}
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	CheckError(err)
	var coral Coral
	//this reports whether another element is available when parsing JSON
	for decoder.More() {
		err := decoder.Decode(&coral)
		CheckError(err)
		corals = append(corals, coral)
	}
	return corals
}

type Coral struct {
	id, types, name, description string
}
