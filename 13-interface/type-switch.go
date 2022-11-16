package main

import "fmt"

func main() {
	//basicSwitch()
	CheckType(true)
}

func CheckType(i any) { // any is an alias to the interface{}
	switch v := i.(type) {
	case int:
		fmt.Println("it is an int", v)
	case string:
		fmt.Println("it is an string", v)

	default:
		fmt.Println("nothing matches")

	}
}

func basicSwitch() {

	day := "mon"

	switch day {

	case "mon":
		fmt.Println("monday")
		fallthrough

	case "tues":
		fmt.Println("tuesday")
		fallthrough

	default:
		fmt.Println("nothing matches")

	}

}
