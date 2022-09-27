package main

import (
	"fmt"
)

type User struct {
	Nombre      string
	Apellido    string
	Edad        int
	Correo      string
	Contrasenha string
}

func cambiarNombre(nombre, apellido string, user *User) {
	user.Nombre = nombre
	user.Apellido = apellido
}

func cambiarEdad(edad int, user *User) {
	user.Edad = edad
}

func cambiarCorreo(correo string, user *User) {
	user.Correo = correo
}

func cambiarContrasenha(contrasenha string, user *User) {
	user.Contrasenha = contrasenha
}

func main() {
	user := User{
		Nombre:      "Jeisson",
		Apellido:    "Santiesteban",
		Edad:        23,
		Correo:      "jsantiesteban99@gmail.com",
		Contrasenha: "jeisson123",
	}
	fmt.Println(user)
	cambiarNombre("fernando", "mendivelso", &user)
	fmt.Println(user)
	cambiarEdad(24, &user)
	fmt.Println(user)
	cambiarCorreo("jeisson.santiesteban@uptc.edu.co", &user)
	fmt.Println(user)
	cambiarContrasenha("jeisson456", &user)
}
