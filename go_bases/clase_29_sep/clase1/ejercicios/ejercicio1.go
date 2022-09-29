package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("./customers.txt")

	defer func() {
		recovered := recover()
		if recovered != nil {
			fmt.Println("ejecución finalizada")
		}
	}()

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}

}
