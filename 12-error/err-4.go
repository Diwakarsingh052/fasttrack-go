package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	var nErr *strconv.NumError
	_, err := strconv.Atoi("efg")
	if errors.As(err, &nErr) {
		fmt.Println(nErr.Err)
	}

	_, err = strconv.ParseInt("1000000", 10, 8)
	fmt.Println(err)

	_, err = strconv.Atoi("abc")
	fmt.Println(err)

}
