package main

import (
	"fmt"
)

func calculator(num1 float64, num2 float64) (float64, float64, float64, float64) {
	suma := num1 + num2
	resta := num1 + num2
	multiplicacion := num1 + num2
	var division float64
	if num2 != 0 {
		division = num1 / num2
	}
	return suma, resta, multiplicacion, division
}

func main() {
	suma, resta, multiplicacion, division := calculator(10, 2)

	fmt.Println("La suma es: ", suma)
	fmt.Println("La resta es: ", resta)
	fmt.Println("La multiplicacion es: ", multiplicacion)
	fmt.Println("La division es: ", division)
}
