package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of Pointers!")

	//declare and create pointer:
	myNumber := 23
	var ptr *int
	ptr = &myNumber

	fmt.Println("Value of Pointer is: ", ptr)
	fmt.Println("Value at Pointer variable is: ", *ptr)

	//string pointer:
	var stringPtr *string
	var str = "Hey!"

	stringPtr = &str
	fmt.Println("Value of pointer (address of variable) is:  ", stringPtr)
	fmt.Println("Value of pointer variable is: ", *stringPtr)

}
