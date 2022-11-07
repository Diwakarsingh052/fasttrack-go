package main

import (
	"fmt"
)

func main() {
	var firstName string
	var abc string // default value is empty string of string datatype
	var i int      // 0

	var name = "Bob"

	age, ok := 30, true

	_, _, _, _ = age, ok, name, firstName // ignore values // don't use in production code

	fmt.Printf("%q\n", abc)
	//fmt.Println(abc)
	fmt.Println(i)

	// in a var block try to put related things together
	var (
		sName  string
		sAge   int = 30
		sMarks float64
	)
	_, _, _ = sName, sAge, sMarks
	//time.Second
	//os.O_TRUNC

}
