package main

import (
	"errors"
	"fmt"
	"os"
)

var CustomError = errors.New("el salario ingresado no alcanza el mínimo imponible")

func main() {
	salary := 0

	if salary < 150000 {
		err := CustomError
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Debe pagar impuesto")
}
