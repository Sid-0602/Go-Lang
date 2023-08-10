package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{"Apple", "Banana", "Guava", "Cherry", "Strawberry"}
	fmt.Println("Fruit List is: ", fruitList)
	fmt.Printf("Type of FruitList is: %T\n", fruitList)
	fmt.Println("Size of Lsit is: ", len(fruitList))

	//append method in list:
	fruitList = append(fruitList, "Jackfruit")
	fmt.Println(" Updated Fruit List is: ", fruitList)
	fmt.Println("Size of Lsit is: ", len(fruitList))

	// [start:end] -> this will include the list from these two positions.
	// start index is included but end is not.
	fruitList = append(fruitList[1:])
	fmt.Println(" Updated Fruit List is: ", fruitList)
	fmt.Println("Size of Lsit is: ", len(fruitList))

	fruitList = append(fruitList[:4])
	fmt.Println(" Updated Fruit List is: ", fruitList)
	fmt.Println("Size of Lsit is: ", len(fruitList))

	fruitList = append(fruitList[1:3])
	fmt.Println(" Updated Fruit List is: ", fruitList)
	fmt.Println("Size of Lsit is: ", len(fruitList))

	highScores := make([]int, 4)
	highScores[0] = 90
	highScores[1] = 95
	highScores[2] = 70
	highScores[3] = 85
	fmt.Println(sort.IntsAreSorted(highScores))

	highScores = append(highScores, 77, 92, 99)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println("Sorted scores are: ", highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value from slice based on index ??

	var courses = []string{"nodejs,", "javascript,", "swift,", "ruby,"}
	courses = append(courses, "react.js,", "sql,", "mongodb")
	fmt.Println("Courses list is: ", courses)

	//remove index from the slice:
	var index = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Course list after index removal is: ", courses)
}
