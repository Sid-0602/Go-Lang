package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	//formatting the time:
	fmt.Println(presentTime.Format("01-02-2006"))

	//create a date:
	createDate := time.Date(2022, time.June, 6, 6, 0, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
