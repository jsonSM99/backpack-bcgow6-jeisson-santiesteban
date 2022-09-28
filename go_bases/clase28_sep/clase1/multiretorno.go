package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrCannotDivideByZero   = errors.New("no se puede dividir por cero")
	ErrOperatorNotSupported = errors.New("Operacion no soportada")
)

func division(num1, num2 float64) (result float64, err error) {
	if num2 == 0 {
		err = ErrCannotDivideByZero
		return
	}
	result = num1 / num2
	return
}

func calcular(operator string, num1, num2 float64) (result float64, err error) {
	switch operator {
	case "/":
		result, err = division(num1, num2)
	default:
		err = ErrOperatorNotSupported
	}
	return
}
func main() {
	result, err := calcular("/", 10, 0)

	if err != nil {
		switch err {
		case ErrCannotDivideByZero:
			fmt.Printf("Ha ocurrido un problema en la capa de división: %s\n", err.Error())
		case ErrOperatorNotSupported:
			fmt.Printf("Ha ocurrido un problema en la capa de calculo: %s\n", err.Error())
		}
		os.Exit(1)
	}

	fmt.Printf("El resultado de la division es: %v\n", result)
}
