package main

import "fmt"

func main() {

	//var i []int // nil

	//pre allocating the underlying array
	i := make([]int, 0, 1000)
	//i = append(i, 100)

	if i == nil {
		fmt.Println("it is nil slice")
	} else {
		fmt.Println("it is not")
	}

	//i = make([]int, 10)
	//for x := 10; x <= 20; x++ {
	//	i[x] = somevalue
	//}
	inspectSlice("i", i)

}

func inspectSlice(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Println(slice)
	fmt.Println()
}
