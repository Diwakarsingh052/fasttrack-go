package main

import "fmt"

func main() {
	var keys []string
	dictionary := map[string]string{"up": "above"} // key : values
	keys = append(keys, "up")
	dictionary["abc"] = "xyz"
	keys = append(keys, "abc")

	dictionary["1"] = "anything"
	dictionary["2"] = "anything"
	dictionary["3"] = "anything"
	dictionary["4"] = "anything"
	dictionary["5"] = "anything"
	fmt.Println(dictionary["abc"])

	// when ranging over maps values would not be returned in the order
	//for k, v := range dictionary {
	//	fmt.Println(k, v)
	//}

	for _, v := range keys {
		fmt.Println(dictionary[v])
	}
}
