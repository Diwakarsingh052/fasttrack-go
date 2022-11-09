package main

import "fmt"

/*
		a:= 10,20,30,40
		b := i[2:5]
		a ,b-> [10,20,(30,40,50),60] // backing array // a and b slice shares the same backing array. it is not a copy
		b[0] = 999 // b is sharing the same backing array with so any changes in b will also result changes in a slice
	    a ,b-> [10,20,(999,40,50),60]
*/
func main() {

	i := []int{10, 20, 30, 40, 50, 60, 70}
	b := i[2:5] // index:len // this is not copying // this is slicing (referencing)
	b[0] = 999  // it will not allocate a new memory so it will change in the base slice as well
	fmt.Println(b)
	fmt.Println(i)

	//fmt.Println("len", len(a), "cap", cap(a))
}
