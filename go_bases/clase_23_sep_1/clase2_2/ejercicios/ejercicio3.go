package main

import "fmt"

func main(){
	numero_mes := 9
	meses := map[int]string {
		1: 	"enero",
		2: 	"febrero",
		3: 	"marzo",
		4: 	"abril",
		5: 	"mayo",
		6: 	"junio",
		7: 	"julio",
		8: 	"agosto",
		9: 	"septiembre",
		10: "octubre",
		11: "noviembre",
		12: "diciembre",
	}
	nombre_mes, existe_mes := meses[numero_mes]
	if !existe_mes {
		fmt.Println("No existe el mes ", numero_mes)
	} else {
		fmt.Println("Mes: ", nombre_mes)
	}
}
/*
	Tambien se puede solucionar con ifs y elseifs
		ejemplo: 
			if(mes == 1) {
				fmt.Println("Mes: enero")
				...
		Nota: se vuelve mucho codigo
	
	Tambien se puede solucionar con switch
		ejemplo:
			switch mes {
				case 1
					fmt.Println("Mes: enero")
		Nota: codigo mas ordenado
	
	NOTA GENERAL: con el uso de maps se puede reducir el codigo
*/