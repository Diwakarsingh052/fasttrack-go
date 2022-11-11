package main

import (
	"fmt"
	"learn-go/db"
	"log"
)

func main() {
	c, err := db.NewConfig("postgres")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(c)

	c1, err := db.NewConfig("mysql")
	fmt.Println(c, c1)

}
