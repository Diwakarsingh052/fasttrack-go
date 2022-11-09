package data

import (
	"fmt"
)

// DBConn // it is a bad idea to export a var in prod especially when things can go wrong when someone changes that var value
var DBConn = "postgres"

var json = "any json"

func GetData() {
	fmt.Println("my prod data is in postgres")
	fmt.Println("getting the data from", DBConn, json)
	//os.Args //
}
