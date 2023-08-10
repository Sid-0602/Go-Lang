package main

import "fmt"

func main() {

	fmt.Println("If-else in Golang")

	//this is if-else control flow for golang:
	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out!"
	} else {
		result = "Exactly 10 login count!"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even!")
	} else {
		fmt.Println("Number is odd!")
	}

	//this is for

	if num := 11; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

}