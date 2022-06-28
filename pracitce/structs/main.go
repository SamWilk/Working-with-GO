package main

import "fmt"

func main() {
	kirby := Dog{"Wired Haired Pointing Griffon", 13, 43, "Woof"}
	dogMap := make(map[string]string)
	dogMap[kirby.breed] = "Kirby"
	//Ok is a boolean flag that gets returned from a search in a map given a key
	//Also you can define variables like so before the if is evaulated
	if value, ok := dogMap["Wired Haired Pointing"]; !ok {
		fmt.Println("No Value Found")
	} else {
		fmt.Println(value)
	}
	fmt.Printf("%+v\n", kirby)
	kirby.SpeakOnce()
	kirby.Speak()
	kirby.Speak()
}

//Dog is a struct
type Dog struct {
	breed  string
	age    int
	weight int
	Sound  string
}

//Speak is how the dog speaks
/*
This is acting in the way of if an "object" of type Dog calls this method
then the method will return that instance (copy of object not a reference) of whatever the field is assigned too
Example: kirby.Sound = "Woof", when this method is called from the kirby variable,
it will print "Woof"
*/
func (d *Dog) Speak() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

//SpeakOnce is how the dog speaks
//Passed a copy here of type Dog
func (d Dog) SpeakOnce() {
	fmt.Println(d.Sound)
}
