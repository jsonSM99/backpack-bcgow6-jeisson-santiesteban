package main

import (
	"errors"
	"fmt"
)

func dividir(dividendo, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("El divisor no puede ser 0")
	}
	return dividendo / divisor, nil
}

func main() {
	result, err := dividir(10, 0)

	if err != nil {
		panic(err) // no usar
	}
	fmt.Println(result)
}
