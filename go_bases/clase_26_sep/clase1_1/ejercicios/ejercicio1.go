package main

import "fmt"

func descontar_impuestos(salario_bruto float64) float64 {
	if salario_bruto > 150000 {
		return salario_bruto * .27
	} else if salario_bruto > 50000 {
		return salario_bruto * .17
	}
	return 0
}

func main() {
	var salario_bruto float64 = 41000

	salario_descontado := descontar_impuestos(salario_bruto)

	fmt.Println("Salario: ", salario_bruto, "Descontado: ", salario_descontado)
}
