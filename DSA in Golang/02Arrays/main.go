package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays in Golang!")

	//array declaration and specification:

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Guava"
	fruitList[2] = "Banana"
	fruitList[3] = "Mango"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit length is: ", len(fruitList))

	//direct declaration:
	var vegList = [5]string{"potato", "beans", "mushrooms"}
	fmt.Println("Veg list is: ", vegList)
	fmt.Println("Veg list size is: ", len(vegList))

}
