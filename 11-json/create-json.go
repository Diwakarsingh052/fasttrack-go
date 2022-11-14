package main

import (
	"encoding/json"
	"log"
	"os"
)

type Perms map[string]bool
type person struct {
	FirstName string                   `json:"first_name"` //field tags // set the field name to the name specified for the json
	Password  string                   `json:"-"`          //- means ignore this field in json output
	Perms     `json:"perms,omitempty"` // if the field is null then avoid in json output
}

func main() {

	users := []person{

		{
			FirstName: "Roy",
			Password:  "abc",
			Perms:     Perms{"admin": true},
		},

		{
			FirstName: "Ray",
			Password:  "qwe",
			Perms:     Perms{"write": false},
		},

		{
			FirstName: "John",
			Password:  "rty",
		},
	}

	//jsonData, err := json.Marshal(users)
	jsonData, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}

	f, err := os.OpenFile("test.json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = f.Write(jsonData)
	if err != nil {
		log.Println(err)
		return
	}

}
