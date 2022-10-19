package calculadora

import "fmt"

func Add(a, b int) (int, error) {
	if a == 0 {
		return 0, fmt.Errorf("a no puede ser %d", 0)
	}
	return a + b, nil
}
func Subs(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, fmt.Errorf("no se aceptan ceros")
	}
	return a - b, nil
}
