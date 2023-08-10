package main

import "fmt"

func main() {
	fmt.Println("This is loops in jump!")

	daysofweek := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	//for loop:
	fmt.Println("Days of weeks are: ")
	for i := 0; i < len(daysofweek); i++ {
		fmt.Println(daysofweek[i])
	}

	/*
		fmt.Println("Days of weeks are: ")
		for i := range daysofweek {
			fmt.Println(daysofweek[i])
		}
	*/

	for _, day := range daysofweek {
		fmt.Printf("Day is %v\n", day)
	}

	//this is while loop and break/continue:
	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 7 {
			goto jump
		}
		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

	//go-to statements:
jump:
	fmt.Println("go-to statement executed at 7.")

}
