package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Args)
	args := os.Args[1:] // start from the first index till the end
	fmt.Println(args)
}
