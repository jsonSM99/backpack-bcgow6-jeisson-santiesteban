package main

import "fmt"
func main()  {
	x,y := 10, 20
	
	fmt.Printf("x + y = %d\n", x + y)
	fmt.Printf("x - y = %d\n", x - y)
	fmt.Printf("x * y = %d\n", x * y)
	fmt.Printf("x / y = %d\n", x / y)
	fmt.Printf("x mod y = %d\n", x % y)
	
	
	fmt.Printf("x++ = %d\n", x++)
	
	fmt.Printf("y-- = %d\n", y--)
}