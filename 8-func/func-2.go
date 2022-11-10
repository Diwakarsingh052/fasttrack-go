package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	res, err := SumString("20", "abc")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res)
}

func SumString(s, x string) (int, error) { // err must be the last value to be returned

	a, err := strconv.Atoi(s)
	if err != nil {

		return 0, err // when err happens then set other values to their defaults
	}

	b, err := strconv.Atoi(x)
	if err != nil {

		return 0, err
	}

	sum := a + b
	return sum, nil // nil indicates no error
}
