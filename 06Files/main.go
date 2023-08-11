package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang!")

	content := "This needs to go in a file- sidbnjadhav@gmail.com. This data is written at time of creating the file!!"

	file, err := os.Create("./myfile.txt") //creating the file in the same folder, using OS packg.

	checkNillErr(err)

	length, err := io.WriteString(file, content) //this writes the given string and returns the length of string.
	checkNillErr(err)

	fmt.Println("Lenght is: ", length)
	defer file.Close() //defer at end close the file.

	//function call to read from file:
	readFile("../06Files/myfile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)

	checkNillErr(err)
	fmt.Println("Text data inside the file is: \n", databyte)
	fmt.Println("Text data inside file in Text format is: \n", string(databyte))

}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
