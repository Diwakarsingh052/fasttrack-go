package main

import "fmt"

func main() {

	//var i int // we can't redeclare var or change there types in a block
	//var i string

	a, b := 10, 20       // creates and assign values
	b, c := 200, "hello" // b gets updated and c var is created
	b = 300
	fmt.Println(a, b, c)

	{
		var b string = "bob"
		fmt.Println(b)
	}

}
