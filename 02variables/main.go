package main

import "fmt"

func main() {

	fmt.Println("This is Variables Section")

	//variables:
	var username string = "Sid"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)
	//bool value:
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("isLoggedIn is of type : %T\n", isLoggedIn)

	//uint and int:

	/*
		uint8 min - max: 0 - 255
		uint16 min - max: 0 - 65535
		uint32 min - max: 0 - 4294967295
		uint64 min - max: 0 - 18446744073709551615
	*/
	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("smallval is of type : %T\n", smallval)

	var smallval2 uint16 = 6550
	fmt.Println(smallval2)
	fmt.Printf("smallval2 is of type : %T\n", smallval2)

	// float
	/*

		float32	---> 32 bits --->-3.4e+38 to 3.4e+38.
		float64	---> 64 bits ---> -1.7e+308 to +1.7e+308.

	*/
	var smallfloat float32 = 255
	fmt.Println(smallfloat)
	fmt.Printf("smallfloat is of type : %T\n", smallfloat)

	var smallfloat2 float64 = 10000
	fmt.Println(smallfloat2)
	fmt.Printf("smallfloat2 is of type : %T\n", smallfloat2)
}
