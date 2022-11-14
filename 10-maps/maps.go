package main

import "fmt"

func main() {

	var dictionary map[string]string // default is nil
	dictionary = make(map[string]string)
	dictionary["up"] = "above"
	dictionary["down"] = "below"

	fmt.Println(dictionary)
	fmt.Println(dictionary["up"])

	// deleting a key
}
