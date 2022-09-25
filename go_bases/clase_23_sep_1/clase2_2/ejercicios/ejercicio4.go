package main

import "fmt"

func main(){
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var num_adults int

	// add employee
	employees["Federico"] = 25
	// delete
	delete(employees, "Pedro")
	
	for _, age := range employees {
		if age > 21 {
			num_adults++
		}
	}
	fmt.Println("Numero de mayores a 21 --> ", num_adults)
}