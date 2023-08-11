package main

import "fmt"

//multiple defer Statements works is LIFO manner.
func main() {

	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("This is defer concept in Go")
	fmt.Println("Hello")

	defer fmt.Println("End")
	defer fmt.Println("World")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

}

//stack of defer: 4,3,2,1,0, world, end, two , one
