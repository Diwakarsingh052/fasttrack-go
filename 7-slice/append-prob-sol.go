package main

import "fmt"

func main() {
	// allocation happens while adding an elm when len == cap or elems that we're trying to add are greater than cap availabe

	i := []int{10, 20, 30, 40, 50, 60, 70}

	b := i[1:5:5] // cap = max-index // we use it to make our length and cap equal
	b[0] = 999    // this will affect our base slice as well because update ops is done before append
	b = append(b, 888)

	inspectSlice("i", i)
	inspectSlice("b", b)

}

func inspectSlice(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Println(slice)
	fmt.Println()
}
