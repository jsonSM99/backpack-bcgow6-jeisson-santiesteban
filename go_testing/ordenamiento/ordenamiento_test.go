package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenarTesting(t *testing.T) {
	// Arrange
	nums := []int{1, 4, 5, 3, 2}
	esperado := []int{1, 2, 3, 4, 5}
	// Act
	res := ordenar(nums)
	// Assert
	if len(res) != len(esperado) {
		t.Errorf("Longitudes diferentes")
	}
	for i, n := range nums {
		if n != esperado[i] {
			t.Errorf("Resultado[%d] = %d es diferente de Esperado[%d] = %d", i, res[i], i, esperado[i])
		}
	}
}
func TestOrdenar(t *testing.T) {
	// Arrange
	nums := []int{1, 4, 5, 3, 2}
	esperado := []int{1, 2, 3, 4, 5}
	// Act
	res := ordenar(nums)
	// Assert
	assert.Lenf(t, res, len(esperado), "La longitud de %v es diferente a %v", res, esperado)
	assert.Equal(t, esperado, res, "Resultado %v es diferente a %v", res, esperado)
}
