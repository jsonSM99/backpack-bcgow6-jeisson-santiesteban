package main

import (
	"fmt"
	"os"
)

func main() {
	salary := 0

	if salary < 150000 {
		err := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Debe pagar impuesto")
}
