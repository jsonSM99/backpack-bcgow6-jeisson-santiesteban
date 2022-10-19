package division

import "fmt"

func division(num, den int) (int, error) {
	if den == 0 {
		return 0, fmt.Errorf("El denominador no puede ser 0")
	}
	return num / den, nil
}
