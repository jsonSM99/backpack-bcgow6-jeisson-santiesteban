package main

import "fmt"

// import "fmt"
const (
	s = "Pequeño"
	m = "Mediano"
	l = "Grande"
)

type TypeOfProducto struct {
	name string
}
type Producto interface {
	CalcularCosto() float64
}
type producto struct {
	tipo   TypeOfProducto
	nombre string
	precio float64
}

func (p producto) CalcularCosto() float64 {
	switch p.tipo.name {
	case s:
		return p.precio
	case m:
		return p.precio + p.precio*.03
	case l:
		return p.precio + (p.precio * .06) + 2500
	}
	return 0
}

type tienda struct {
	productos []Producto
}

func nuevoProducto(tipo, nombre string, precio float64) Producto {
	return &producto{
		tipo: TypeOfProducto{
			tipo,
		},
		nombre: nombre,
		precio: precio,
	}
}

func nuevaTienda() Ecommerce {
	return &tienda{}
}

type Ecommerce interface {
	Total() float64
	Agregar(Producto)
}

func (t tienda) Total() float64 {
	var total float64

	for _, p := range t.productos {
		total += p.CalcularCosto()
	}
	return total
}

func (t *tienda) Agregar(p Producto) {
	t.productos = append(t.productos, p)
}

func main() {
	store := nuevaTienda()

	cama := nuevoProducto(l, "cama", 1500000)
	nevera := nuevoProducto(m, "nevera", 2000000)
	lavadora := nuevoProducto(s, "lavadora", 700000)

	fmt.Println(cama.CalcularCosto())
	fmt.Println(nevera.CalcularCosto())
	fmt.Println(lavadora.CalcularCosto())
	store.Agregar(cama)
	store.Agregar(nevera)
	store.Agregar(lavadora)

	fmt.Printf("Total --> %.2f\n", store.Total())
}
