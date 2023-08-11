package main

import "fmt"

func main() {
	fmt.Println("This are functions in Golang!")
	sayHi()
	Namaste()

	result := adder(3, 4)
	fmt.Println("Result of addition of these numbers: ", result)
	//IIF (Immediately invoked Functions):
	squareof2 := func() int {
		return 2 * 2
	}()

	fmt.Println("The result of 2 square is (calculated using IIF) ", squareof2)

	//this is pro function:

	proResult := proAdder(2, 55, 6, 7, 8, 22, 1, 99)
	fmt.Println("Pro Result is:: ", proResult)
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}

	return total
}
func adder(valOne int, valTwo int) int {
	return (valOne + valTwo)
}

func sayHi() {
	fmt.Println("Say hi function!!")
	Namaste() //calling another function.
}

func Namaste() {
	fmt.Println("Namaste from Golang!!")
}
