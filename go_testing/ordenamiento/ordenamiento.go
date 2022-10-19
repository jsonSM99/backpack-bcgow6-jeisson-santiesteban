package ordenamiento

import "sort"

func ordenar(numeros []int) []int {
	sort.Slice(numeros, func(i, j int) bool {
		return numeros[i] < numeros[j]
	})
	return numeros
}
