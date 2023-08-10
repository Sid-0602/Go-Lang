package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["CPP"] = "Cplusplus"
	languages["GO"] = "Go Lang"

	//get entire value maps:
	fmt.Println("List of all Languages: ", languages)

	//get a single value from map:
	fmt.Println("JS Shorts for: ", languages["JS"])
	fmt.Print("PY Shorts for: ", languages["PY"])

	//delete
	delete(languages, "RB")
	fmt.Println("Language: ", languages)

	//looping through maps:

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	//to ignore anything, we use _ in go:
	for _, value := range languages {
		fmt.Printf(" Languages list:  %v\n", value)
	}

}
