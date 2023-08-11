package main

import "fmt"

func main() {
	fmt.Println("Methods in Golang!")
	fmt.Println("Structs in Golang")

	//no inheritance in golang: No super or parent/child.

	sid := User{"Sid", "sidjadhav@gmail.com", true, 21}
	fmt.Println("The user is: ", sid)
	fmt.Printf("Sid's details are: %+v\n", sid)
	fmt.Printf("Name is %v and email is %v\n", sid.Name, sid.Email)

	sid.GetStatus()
	sid.NewMail()
	fmt.Printf("Name is %v and email is %v\n", sid.Name, sid.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

//method to find and parse user's online status:
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

//method uses a copy of object and not reference.
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
