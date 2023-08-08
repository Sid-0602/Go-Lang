package main

import "fmt"

func main() {

	//initialize variables:

	var smsSendingLimit int
	var costPerSms float64
	var hasPermission bool
	var username string

	//short assignment operator:
	empty := ""
	numCars := 10
	var isFunny = true

	fmt.Println("Hello, GoLang!!")
	fmt.Printf(
		"%v %f %v %q %q\n",
		smsSendingLimit,
		costPerSms,
		hasPermission,
		username,
		empty,
	)

	fmt.Printf(
		"%v %v\n",
		numCars,
		isFunny,
	)

}
