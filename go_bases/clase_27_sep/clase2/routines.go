package main

import (
	"fmt"
	"time"
)

func procesar(i int, ch chan int) {
	fmt.Println(i, "Inicia!")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "Termina")
	ch <- i
}

func main() {
	canal := make(chan int)
	now := time.Now()
	fmt.Println("Inicia el programa")
	for i := 0; i < 10; i++ {
		go procesar(i, canal)
	}
	for i := 0; i < 10; i++ {
		lectura := <-canal
		fmt.Println(lectura)
	}
	fmt.Println("Termina el programa")
	fmt.Println("El programa demoro: ", time.Now().Sub(now).Milliseconds())
}
