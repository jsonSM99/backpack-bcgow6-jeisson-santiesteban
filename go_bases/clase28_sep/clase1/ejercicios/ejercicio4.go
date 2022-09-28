package main

import (
	"errors"
	"fmt"
)

type Salario struct {
	horasTrabajadas int
	valorHora       float64
	mes             int
}
type Trabajador struct {
	nombre   string
	salarios []Salario
}

func CalcularSalario(s Salario) (salario float64, err error) {

	if s.horasTrabajadas < 80 {
		err = errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
		return
	}

	salario = float64(s.horasTrabajadas) * s.valorHora

	if salario >= 150000 {
		salario -= salario * .1
	}
	return
}

func CalcularMedioAguinaldo(trabajador Trabajador) (aguinaldo float64, err error) {
	var mejorSalario float64
	var numeroMesesTrabajados int
	mejorSalario, numeroMesesTrabajados, err = CalcularDatosSalario(trabajador)

	if err != nil {
		return
	}
	aguinaldo = (mejorSalario / 12) * float64(numeroMesesTrabajados)
	return
}

func CalcularDatosSalario(trabajador Trabajador) (mejorSalario float64, numeroMesesTrabajados int, err error) {
	for _, salario := range trabajador.salarios {
		var s float64
		s, err = CalcularSalario(salario)
		if err != nil {
			return
		}
		if s > mejorSalario {
			mejorSalario = s
		}
		numeroMesesTrabajados++
	}
	return
}

func main() {
	trabajador := Trabajador{
		nombre: "jeisson",
		salarios: []Salario{
			{
				horasTrabajadas: 160,
				valorHora:       3000,
				mes:             1,
			},
			{
				horasTrabajadas: 160,
				valorHora:       3000,
				mes:             2,
			},
			{
				horasTrabajadas: 161,
				valorHora:       3000,
				mes:             3,
			},
			{
				horasTrabajadas: 167,
				valorHora:       3000,
				mes:             4,
			},
			{
				horasTrabajadas: 155,
				valorHora:       3000,
				mes:             5,
			},
			{
				horasTrabajadas: 140,
				valorHora:       3000,
				mes:             6,
			},
		},
	}

	salario, err := CalcularSalario(Salario{
		horasTrabajadas: 140,
		valorHora:       3000,
		mes:             6,
	})

	if err != nil {
		fmt.Println("1", err.Error())
	}

	fmt.Println("Salario calculado: ", salario)

	aguinaldo, err := CalcularMedioAguinaldo(trabajador)

	if err != nil {
		fmt.Println("2", err.Error())
	}

	fmt.Println("aguinaldo calculado : ", aguinaldo)
}
