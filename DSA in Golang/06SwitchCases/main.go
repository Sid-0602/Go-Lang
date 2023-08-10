package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("This is Switch and Case in Golang")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input any no. between 1-6: ")

	diceNumber, _ := reader.ReadString('\n')
	fmt.Println("Value of Dice is: ", diceNumber)

	//switch and case statements:

	switch diceNumber {
	case "1":
		fmt.Println("Dice value is 1, you can Open")
	case "2":
		fmt.Println("You can move 2 spots")
	case "3":
		fmt.Println("You can move 3 spots")
	case "4":
		fmt.Println("You can move 4 spots")
	case "5":
		fmt.Println("You can move 5 spots")
	case "6":
		fmt.Println("You can move 6 spots and roll the dice again!")
	default:
		fmt.Println("What was that?? Input value b/w 1-6 only!!")

	}
}
