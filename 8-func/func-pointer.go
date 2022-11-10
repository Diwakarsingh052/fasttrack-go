package main

import "fmt"

func main() {
	i := 30 // i address = x809
	x := &i // x address =  x106
	update(x)
	fmt.Println(i)
	fmt.Println(&x)
}

func update(y *int) { // i address would be copied here , x and y are now two different pointers pointing to the same memory loc
	*y = 20 // y address = x108
	fmt.Println(&y)
}
