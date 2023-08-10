package main

import "fmt"

//structs are equivalent to classes in Golang.

func main() {

	fmt.Println("Structs in Golang")

	//no inheritance in golang: No super or parent/child.

	sid := User{"Sid", "sidjadhav@gmail.com", true, 21}
	fmt.Println("The user is: ", sid)
	fmt.Printf("Sid's details are: %+v\n", sid)
	fmt.Printf("Name is %v and email is %v\n", sid.Name, sid.Email)

}

//defining struct:
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
