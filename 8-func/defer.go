package main

import "fmt"

func main() {
	//defer maintains a stack. value added first would be exec at last
	defer fmt.Println("1") // when your function is returning defer statements will exec  //[2,1]
	defer fmt.Println("2")
	fmt.Println("3")
	//return
	var i []int
	i[1] = 100
	//panic("i need to panic") // defer guarantee exec // if the statements are registered before the panic or return
	fmt.Println("4")
}
