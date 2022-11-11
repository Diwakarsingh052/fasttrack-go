package main

import (
	"fmt"
)

func main() {
	show("data", 10, 90) // variadic param are optional
	show2([]int{102, 30, 50, 60})
}

func show2(i []int) {}

func show(s string, i ...int) { // variadic param should be the last in the func signature
	fmt.Println(s, i)
	fmt.Printf("%#v\n", i)

}
