package main

import (
	"fmt"
)

type Productos struct {
	nombre   string
	precio   float64
	cantidad int
}

type Servicios struct {
	nombre            string
	precio            float64
	minutosTrabajados int
}
type Mantenimiento struct {
	nombre string
	precio float64
}

func SumarProductos(productos *[]Productos, ch chan float64) float64 {
	var result float64
	for _, producto := range *productos {
		result += producto.precio * float64(producto.cantidad)
	}
	ch <- result
	return result
}

func SumarServicios(servicios *[]Servicios, ch chan float64) float64 {
	var result float64
	for _, servicio := range *servicios {
		media_hora_aux := servicio.minutosTrabajados / 30
		fmt.Println(media_hora_aux)
		var media_hora float64 = float64(media_hora_aux)
		if media_hora_aux < 1 {
			media_hora = 1
		}
		result += servicio.precio * float64(media_hora)
	}
	ch <- result
	return result
}

func SumarMantenimiento(mantenimientos *[]Mantenimiento, ch chan float64) float64 {
	var result float64
	for _, mantenimiento := range *mantenimientos {
		result += mantenimiento.precio
	}
	ch <- result
	return result
}
func main() {
	canalProductos := make(chan float64)
	canalServicios := make(chan float64)
	canalMantenimiento := make(chan float64)

	productos := []Productos{
		{
			nombre:   "lavadora",
			precio:   1990000,
			cantidad: 1,
		},
		{
			nombre:   "microhondas",
			precio:   200000,
			cantidad: 2,
		},
		{
			nombre:   "tv",
			precio:   1700000,
			cantidad: 3,
		},
	}
	servicios := []Servicios{
		{
			nombre:            "lavar",
			precio:            50000,
			minutosTrabajados: 31,
		},
		{
			nombre:            "cocinar",
			precio:            40000,
			minutosTrabajados: 64,
		},
		{
			nombre:            "organizar",
			precio:            200000,
			minutosTrabajados: 83,
		},
	}
	mantenimientos := []Mantenimiento{
		{
			nombre: "carro",
			precio: 3000000,
		},
		{
			nombre: "moto",
			precio: 1220000,
		},
		{
			nombre: "casa",
			precio: 800000,
		},
	}
	go SumarProductos(&productos, canalProductos)
	go SumarServicios(&servicios, canalServicios)
	go SumarMantenimiento(&mantenimientos, canalMantenimiento)

	lecturaProductos := <-canalProductos
	lecturaServicios := <-canalServicios
	lecturaMantenimiento := <-canalMantenimiento

	fmt.Printf("lecturaProductos --> %.0f\n", lecturaProductos)
	fmt.Printf("lecturaServicios --> %.0f\n", lecturaServicios)
	fmt.Printf("lecturaMantenimiento --> %.0f\n", lecturaMantenimiento)
	fmt.Printf("Suma total --> %.0f \n", lecturaProductos+lecturaServicios+lecturaMantenimiento)
}
