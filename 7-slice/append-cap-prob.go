package main

import "fmt"

/*
   note () indicates b elements
	a:= 10,20,30,40,50, 60, 70
	b := a[2:5] // slicing
	a ,b-> [10,20,(30,40,50,888),70,80] // backing array

	b = append(b, 888)

	a,b ->[10 20 30 40 50 888 70]
*/

// cap is calc for a slice from the base index where b slice starts and till the end of the underlying array
func main() {
	i := []int{10, 20, 30, 40, 50, 60, 70, 80}
	b := i[2:5]

	//b = append(b, 888) //It is going to change the base slice because they are  sharing the same memory
	inspectSlice("b", b)
	b = append(b, 888, 999, 777, 666) // we are adding to add 4 elems here but we don't have enough cap to fit all the values, allocation will happen here
	inspectSlice("i", i)
	inspectSlice("b", b)
}

func inspectSlice(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Println(slice)
	fmt.Println()
}
