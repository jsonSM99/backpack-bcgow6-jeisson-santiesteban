package main

import (
	"fmt"
	"os"
)

func main() {
	data := `Id,Precio,Cantidad
111223,30012.00,1
444321,1000000.00,4
434321,50.50,1`

	err := os.WriteFile(
		"./products.csv",
		[]byte(data),
		0644,
	)

	if err != nil {
		panic(err)
	}
	fmt.Println("Archivo creado!")
}
