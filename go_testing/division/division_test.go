package division

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivisionTesting(t *testing.T) {
	// Arrange
	num1 := 1
	num2 := 0
	// Act
	result, err := division(num1, num2)
	// Assert
	if err == nil || result != 0 {
		t.Errorf("Se ejecutó una division por cero")
	}
}
func TestDivision(t *testing.T) {
	// Arrange
	num1 := 1
	num2 := 0
	errorEsperado := "El denominador no puede ser 0"
	// Act
	result, err := division(num1, num2)
	// Assert
	assert.ErrorContains(t, err, errorEsperado)
	assert.Equal(t, 0, result)
}
