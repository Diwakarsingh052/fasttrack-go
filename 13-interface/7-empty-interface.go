package main

import "fmt"

type stu struct{}

func main() {
	var s stu
	var i interface{}
	x := 100
	check := true
	i = x
	i = check
	i = s

	y, ok := i.(int) // type assertion
	if !ok {
		fmt.Println("not int")
		return
	}
	fmt.Println(y)

}

// create a variadic func that accepts two values of any type but if integers are passed than do addition with them
