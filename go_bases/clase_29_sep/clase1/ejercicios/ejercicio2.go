package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

type Client struct {
	Docket    int
	FirstName string
	Lastname  string
	DNI       string
	Address   string
}

var (
	IdNotGenerated = "Id no generado"
	FileNotFound   = "el archivo indicado no fue encontrado o está dañado"
)

func generateID() uuid.UUID {
	return uuid.New()
}
func CheckClient(clientToCheck Client) (client Client, err error) {

	defer func() {
		if recovered := recover(); recovered != nil {
			err = recovered.(error)
		}
	}()

	if IsZero(clientToCheck.Docket) {
		err = errors.New("Docket esta vacio")
	}
	if IsZero(clientToCheck.FirstName) {
		err = errors.New("FirstName esta vacio")
	}
	if IsZero(clientToCheck.Lastname) {
		err = errors.New("Lastname esta vacio")
	}
	if IsZero(clientToCheck.DNI) {
		err = errors.New("DNI esta vacio")
	}
	if IsZero(clientToCheck.Address) {
		err = errors.New("Address esta vacio")
	}
	if err != nil {
		panic(err)
	}
	client = clientToCheck
	return
}

func IsZero[T comparable](v T) bool {
	fmt.Println("XD", reflect.TypeOf(v), v == *new(T))
	return v == *new(T)
}

func ReadFile(address string) (customers *os.File, err error) {
	// prevenir el cierre inesperado
	defer func() {
		if recovered := recover(); recovered != nil {
			customers.Close()
			err = fmt.Errorf("No han quedado archivos abiertos: %w", err)
			err = fmt.Errorf("Se detectaron varios errores en tiempo de ejecución: %w", err)
			err = fmt.Errorf("Fin de la ejecución: %w", err)
			err = fmt.Errorf(recovered.(error).Error()+": %w", err)
		}

	}()
	// leer datos de clientes ya existentes
	customers, err = os.Open(address)

	if err != nil {
		panic(errors.New(FileNotFound))
	}
	return
}

func main() {

	customers, err := ReadFile("./customers.txt")

	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(customers)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	customers.Close()

	for _, line := range text {
		fields := strings.Split(line, ";")
		_, err = CheckClient(Client{
			Docket:    os.Geteuid(),
			FirstName: fields[0],
			Lastname:  fields[1],
			DNI:       fields[2],
			Address:   fields[3]})
		fmt.Println(err)
	}
}
