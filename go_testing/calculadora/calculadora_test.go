package calculadora

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBad(t *testing.T) {
	//Arrange
	num1 := 0
	num2 := 5
	errorEsperado := fmt.Sprintf("a no puede ser %d", num1)
	//Act
	_, err := Add(num1, num2)
	//Assert
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, errorEsperado)
}

func TestAddGood(t *testing.T) {
	//Arrange
	num1 := 10
	num2 := 5
	esperado := 15
	//Act
	resultado, err := Add(num1, num2)
	//Assert
	assert.Equal(t, esperado, resultado, "El numero resultado: %d es distinto del esperado: %d", resultado, esperado)
	assert.Nil(t, err)
	// if resultado != esperado {
	// 	t.Errorf("El numero resultado: %d es distinto del esperado: %d", resultado, esperado)
	// }
}

func TestSubBad(t *testing.T) {
	// Arrange
	num1 := 0
	num2 := 2
	errorEsperado := "no se aceptan ceros"
	// Act
	_, err := Subs(num1, num2)
	// Assert
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, errorEsperado)
}
func TestSubBadTesting(t *testing.T) {
	// Arrange
	num1 := 0
	num2 := 2
	errorEsperado := fmt.Errorf("no se aceptan ceros")
	// Act
	_, err := Subs(num1, num2)
	// Assert
	if err != nil && errors.Is(err, errorEsperado) {
		t.Errorf("numeros con ceros")
	}
}
func TestSubGoodTesting(t *testing.T) {
	// Arrange
	num1 := 1
	num2 := 2
	esperado := -1
	// Act
	resultado, err := Subs(num1, num2)
	// Assert
	if err != nil && esperado != resultado {
		t.Errorf("el número esperado %d es diferente del resultado %d", esperado, resultado)
	}
}
