package main

import (
	"fmt"
	"strings"
)

func main(){
	var clientes 					= map[string]map[string]int{}
	clientes["jeisson"] 				= make(map[string]int);
	clientes["jeisson2"] 				= make(map[string]int);
	clientes["jeisson3"] 				= make(map[string]int);
	clientes["jeisson"]["edad"] 			= 21
	clientes["jeisson"]["antiguedad"] 		= 0
	clientes["jeisson"]["sueldo"] 			= 24000
	clientes["jeisson2"]["edad"] 			= 22
	clientes["jeisson2"]["antiguedad"] 		= 2
	clientes["jeisson2"]["sueldo"] 			= 70000
	clientes["jeisson3"]["edad"] 			= 27
	clientes["jeisson3"]["antiguedad"] 		= 3
	clientes["jeisson3"]["sueldo"] 			= 124000

	for key, cliente := range clientes {
		var str strings.Builder
		str.WriteString("Cliente: " + key)
		if (cliente["edad"] < 22) || (cliente["antiguedad"] < 1){
			fmt.Println(str.String() + " DESAPROBADO")
			continue
		}
		if cliente["sueldo"] < 100000 {
			str.WriteString(" Aplica interes")
		}
		fmt.Println(str.String() +  " APROBADO")
    }

}