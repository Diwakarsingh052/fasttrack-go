package main

import "fmt"

//arrays size is fixed

func main() {
	var a [5]int
	var b = [5]int{10, 20, 30}
	c := [5]int{10, 20, 30}
	d := [...]int{10, 40, 400, 60, 50}

	//d[6] = 888 // we cant grow arrays

	fmt.Println(a)
	fmt.Println(b, c, d)
	fmt.Println(len(d))

	for i, v := range d { // i = index and v = values
		fmt.Println(i, v)
	}

	for _, v := range d { // _ = ignore values and v = values
		fmt.Println(v)
	}

	for i := range d {
		fmt.Println(i)
	}
}
