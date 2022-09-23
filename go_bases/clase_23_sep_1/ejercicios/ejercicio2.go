package main

import (
	"fmt"
)

func main() {
	var temperatura float64 // numeros grandes con decimales
	var humedad int // se mide en porcentaje
	var presion float64 // numeros grandes con decimales

	temperatura = 18
	humedad = 65
	presion = 1013.21

	fmt.Println("Temperatura:",temperatura, "C")
	fmt.Println("Humedad:",humedad,"%")
	fmt.Println("Presión:",presion,"mbar")
}