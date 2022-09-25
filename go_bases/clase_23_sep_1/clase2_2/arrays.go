package main

import "fmt"

func main()  {
	var myArray = [...]int{1,2,3,4,5}
	
	var slc2 = []int

	slc := myArray[0:2]

	slc = append(slc, 5,6,7,7,8,9,5,6,7,7,8,9,5,6,7,7,8,9)

	fmt.Println(cap(slc), len(slc))
}
