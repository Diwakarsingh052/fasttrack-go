package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type user struct {
	Name  string          `json:"first_name"`
	Roles map[string]bool `json:"perms"`
}

func main() {
	jsonData, err := os.ReadFile("test.json")
	if err != nil {
		log.Fatalln(err)
	}
	var u []user
	err = json.Unmarshal(jsonData, &u)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(u)

}
