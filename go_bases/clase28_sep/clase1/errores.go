package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// err1 := fmt.Errorf("first error :(")

	// err2 := fmt.Errorf("second error: %w", err1)

	// err3 := errors.New("third error :(")

	// fmt.Println(err1.Error())
	// fmt.Println(err2.Error())
	// fmt.Println(err3.Error())

	// statusCode := 404

	// if statusCode > 399 {
	// 	err := errors.New("La peticion ha fallado")
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("Status %d, funciona! :D", statusCode)

	// err1 := &MyCustomError{
	// 	StatusCode: 502,
	// 	Message:    "Algo salió mal :( 1",
	// }
	// err2 := &MyCustomError{
	// 	StatusCode: 406,
	// 	Message:    "Algo salió mal :( 2",
	// }

	// bothErrorsAreEqual := errors.As(err1, &err2)

	// fmt.Println("Son iguales --> ", bothErrorsAreEqual)

	// _, err := find()

	// if err != nil {
	// 	if errors.Is(err, ErrNotFound) {
	// 		fmt.Println("No encontrado")
	// 	} else {
	// 		fmt.Println("Error interno")
	// 	}
	// }

	err2 := ErrorType2{}
	err1 := fmt.Errorf("soy el error 1, contengo a %w", err2)
	err := fmt.Errorf("soy el error, contengo a %w", err1)

	fmt.Println(
		errors.Unwrap(err),
	)

	fmt.Println(
		errors.Unwrap(err1),
	)

	fmt.Println(
		errors.Unwrap(err2),
	)
}

// //////
type ErrorType2 struct {
}

func (e ErrorType2) Error() string {
	return "soy el error 2"
}

// /////
var ErrNotFound = errors.New("not found")

func find() (result string, err error) {
	return "", ErrNotFound
}

// ////
type MyCustomError struct {
	StatusCode int
	Message    string
}

func (err *MyCustomError) Error() string {
	return fmt.Sprintf(
		"%s (%d)",
		err.Message,
		err.StatusCode,
	)
}

func ObtainError(status int) (code int, err error) {
	if status >= 500 {
		return 500, &MyCustomError{
			StatusCode: code,
			Message:    "Algo salió mal :(",
		}
	}
	return 200, nil
}
func CustomErrors() {
	status, err := ObtainError(500)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Status %d, funciona! :D", status)
}
