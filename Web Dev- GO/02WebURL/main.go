package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&author=lco"

func main() {
	fmt.Println("Handling URLs in Golang!")
	fmt.Println(myurl)

	//parsing:

	result, _ := url.Parse(myurl)

	//this gives back url components:
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of query params are: %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	//constructing a URL:
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=sid",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
