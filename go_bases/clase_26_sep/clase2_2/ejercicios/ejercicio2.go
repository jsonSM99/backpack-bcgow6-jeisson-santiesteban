package main

import "fmt"

// import "fmt"
type Matrix struct {
	val [][]int
}

func (m Matrix) getMatrixInfo() (is_cuadratica bool, x_dim int, y_dim int) {
	x_dim = len(m.val)

	if x_dim <= 0 {
		return false, x_dim, y_dim
	}

	y_dim = len(m.val[0])

	if y_dim <= 0 {
		return false, x_dim, y_dim
	}
	return x_dim == y_dim, x_dim, y_dim
}

func (m Matrix) Print() {
	for _, x := range m.val {
		for _, y := range x {
			fmt.Printf("%v ", y)
		}
		fmt.Printf("\n")
	}
}

func (m *Matrix) Set(new [][]int) {
	m.val = new
}

func main() {

	matrix := Matrix{
		val: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	}
	matrix.Print()

	matrix.Set([][]int{
		{1, 2, 3},
		{7, 8, 9},
		{4, 5, 6},
	})

	matrix.Print()

	is_cuadratica, x_dim, y_dim := matrix.getMatrixInfo()
	fmt.Printf("Cuadratica: %t\tDimension X: %v\tDimension Y: %v\n", is_cuadratica, x_dim, y_dim)
}
