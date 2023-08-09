package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	//BUFIO AND OS PCKG for user-input:
	reader := bufio.NewReader(os.Stdin) //bufio is pckg
	fmt.Println("Enter the rating for service: ")

	// comma ok || err ok:

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us with : ", input)
	fmt.Printf("Type of rating is: %T \n", input)

}
