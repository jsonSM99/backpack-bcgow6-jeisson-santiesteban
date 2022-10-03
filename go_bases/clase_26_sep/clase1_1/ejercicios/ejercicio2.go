package main

import "fmt"

func calcular_promedio_calificaciones(calificaciones ...float64) float64 {
	var result float64
	n := len(calificaciones)

	if n == 0 {
		return result
	}

	for _, calificiacion := range calificaciones {
		result += calificiacion
	}
	return result / float64(n)
}

func main() {

	promedio_calificiaciones := calcular_promedio_calificaciones(4, 5, 3.3, 3.7, 4.7, 3.9, 4)

	fmt.Println("Promedio de calificaciones:", fmt.Sprintf("%.2f", promedio_calificiaciones))
}
