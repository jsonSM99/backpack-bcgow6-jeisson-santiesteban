package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	products, err := os.Open("./products.csv")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File opened!")
	defer products.Close()

	csvLines, err := csv.NewReader(products).ReadAll()

	if err != nil {
		fmt.Println(err)
	}
	var total float64
	for _, row := range csvLines {
		var precio float64
		var cantidad float64
		for j, col := range row {
			if num, err := strconv.ParseFloat(col, 64); err == nil && j == 1 {
				precio = num
			}
			if num, err := strconv.ParseFloat(col, 64); err == nil && j == 2 {
				cantidad += num
			}
			total += precio * cantidad
			fmt.Print(col + "\t\t")
		}
		fmt.Println()
	}
	fmt.Printf("\t\t%.2f\n", total)
}
