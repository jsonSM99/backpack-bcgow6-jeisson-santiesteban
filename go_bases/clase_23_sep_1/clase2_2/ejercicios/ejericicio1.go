package main

import "fmt"

func main(){
	palabra := "jeisson fernando santiesteban mendivelso"
	fmt.Printf("longitud -->", len(palabra))
	for _, char := range palabra {
		fmt.Printf("Caracter --> %v\n", string(char))
	} 
}