package calc

import "fmt"

var sum int

func Add() {
	i, b := 10, 20
	sum = i + b
	printSomething()
	//fmt.Println("i am add func from the calc package", sum)
}

func printSomething() {
	fmt.Println("i am add func from the calc package", sum)
}
