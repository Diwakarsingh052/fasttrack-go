package main

import "fmt"

type student struct { // student is a new data type
	name  string //fields
	age   int
	marks []float64
}

func main() {
	var u student          // initialize all the fields with its zero value
	fmt.Printf("%#v\n", u) // %#v is for extra info
	fmt.Printf("%+v\n", u) // %+v to print fields and data

	u.name = "bob"
	u.marks = []float64{60.55}
	fmt.Printf("%+v\n", u)

	n := u.name

	fmt.Printf("%T\n", n)

	for i, v := range u.marks {
		fmt.Println(i, v)
	}

}
