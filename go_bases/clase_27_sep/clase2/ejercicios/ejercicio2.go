package main

import (
	"fmt"
)

type Product struct {
	Nombre   string
	precio   float64
	cantidad int
}

type User struct {
	Nombre   string
	Apellido string
	Edad     int
	Correo   string
	Products []Product
}

func NuevoProducto(nombre string, precio float64) Product {
	return Product{
		Nombre: nombre,
		precio: precio,
	}
}

func AgregarProducto(user *User, producto *Product, cantidad int) {
	producto.cantidad = cantidad
	user.Products = append(user.Products, *producto)
}

func BorrarProductos(user *User) {
	user.Products = nil
}
func main() {
	user := User{
		Nombre:   "Jeisson",
		Apellido: "Santiesteban",
		Edad:     23,
		Correo:   "jsantiesteban99@gmail.com",
	}

	lavadora := NuevoProducto("lavadora", 1990000)
	AgregarProducto(&user, &lavadora, 2)
	nevera := NuevoProducto("nevera", 3000000)
	AgregarProducto(&user, &nevera, 2)
	notebook := NuevoProducto("notebook", 12000000)
	AgregarProducto(&user, &notebook, 1)

	fmt.Println(user.Products)

	BorrarProductos(&user)
	fmt.Println(user.Products)
}
