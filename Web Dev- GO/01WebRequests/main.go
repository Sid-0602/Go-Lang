package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web Request in Golang")

	response, err := http.Get(url)

	nilerrhandle(err) //function call to handle error in code.
	fmt.Printf("Response is of type %T\n", response)
	fmt.Printf("******************************************** RESPONSE ******************************************** \n")

	defer response.Body.Close() //close the connection.

	databytes, err := io.ReadAll(response.Body)
	nilerrhandle(err)

	content := string(databytes)
	fmt.Println(content)
}

func nilerrhandle(err error) {
	if err != nil {
		panic(err)
	}
}
