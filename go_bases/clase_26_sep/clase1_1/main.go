package main

const (
	OperadorSuma           = "+"
	OperadorResta          = "-"
	OperadorMultiplicacion = "*"
	OperadorDivision       = "/"
)

// func inspeccionarVariable(numero int) {
// 	if numero > 0 {
// 		fmt.Println("numero positivo")
// 	} else if numero < 0 {
// 		fmt.Println("numero negativo")
// 	} else {
// 		fmt.Println("numero cero")
// 	}
// }

// func suma(num1, num2 float64) float64 {
// 	result := num1 + num2
// 	return result
// }

func operacionAritmetica(num1, num2 float64, operador string) float64 {
	var result float64
	switch operador {
	case OperadorSuma:
		result = num1 + num2
	case OperadorResta:
		result = num1 - num2
	case OperadorMultiplicacion:
		result = num1 * num2
	case OperadorDivision:
		result = num1 / num2
	}
	return result
}
func main() {
	// a, b, c := 0 , 2, -1

	// inspeccionarVariable(a)
	// inspeccionarVariable(b)
	// inspeccionarVariable(c)

	// a := suma(25.5, 3.0)
	// b := suma(23.3, 23.3)

	// fmt.Println(a, b)
}
