package main

import "fmt"

func main() {
	var mapa = map[string]int{"jeisson": 23, "fernando": 25}

	mapa["santi"] = 10

	x, ok := mapa["santi"]
	fmt.Println(x,ok)
}