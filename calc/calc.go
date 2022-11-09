package calc

import "fmt"

var Sum int

// Add func is exported as first letter is uppercase
func Add() {
	i, b := 10, 20
	Sum = i + b
	sub()

	//fmt.Println("i am add func from the calc package", sum)
}

func PrintSomething() {
	fmt.Println("i am add func from the calc package", Sum)
}
