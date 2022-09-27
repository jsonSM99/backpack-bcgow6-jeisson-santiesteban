package main

import "fmt"

func main() {
	messageHello := "Hello world \n"
	messageBye := "Bye world :( \\\\"

	b, err := fmt.Print(messageHello, messageBye)

	if err != nil {
		panic(err)
	}

	fmt.Println(b)
}
