package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) updateEmail(email string) {
	u.email = email
}

type admin struct {
	user // embedding // anonymous field because we are not giving any name to it // anonymous field name would be assigned according to the type name
	role []string
}

func (a admin) show() {
	fmt.Printf("%+v\n", a)
}

func main() {
	var u user = user{
		name:  "a",
		email: "a@email.com",
	}
	_ = u
	// admin is embedding user struct ,so it can use all its fields and its methods
	a := admin{
		user: u,
		role: []string{"admin"},
	}
	fmt.Println(a.name)
	a.updateEmail("bob@gmail.com")
	a.show()
	//accept(a) // it will not work

}

func accept(u user) {

}
