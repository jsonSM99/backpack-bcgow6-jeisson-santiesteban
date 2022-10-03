package main

import "fmt"

//import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (a Alumno) Detalle() {
	fmt.Printf(`
	Nombre: 	%v
	Apellido: 	%v
	DNI: 		%v
	Fecha: 		%v
	
	`,
		a.Nombre,
		a.Apellido,
		a.DNI,
		a.Fecha,
	)
}
func main() {
	alumnos := []Alumno{
		{
			Nombre:   "jeisson",
			Apellido: "santiesteban",
			DNI:      "1115866294",
			Fecha:    "02/06/1999",
		},
		{
			Nombre:   "jeisson2",
			Apellido: "santiesteban2",
			DNI:      "1115866294",
			Fecha:    "02/06/1998",
		},
	}

	for _, alumno := range alumnos {
		alumno.Detalle()
	}
}
