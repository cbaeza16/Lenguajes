package main

import (
	"fmt"
	"sort"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

// la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones
const existenciaMinima int = 10

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	for i := range *l {

		if (*l)[i].nombre == nombre {
			(*l)[i].cantidad += cantidad
			if (*l)[i].precio != precio {
				(*l)[i].precio = precio
			}
			return
		}
	}
	*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
}

// Función adicional que reciba una cantidad potencialmente infinita de productos previamente creados y los agregue a la lista
func (l *listaProductos) agregarProductosInfinita(productos ...producto) {
	for _, nuevoProducto := range productos {
		l.agregarProducto(nuevoProducto.nombre, nuevoProducto.cantidad, nuevoProducto.precio)
	}
}

func (l *listaProductos) buscarProducto(nombre string) (*producto, int) {
	var producto *producto
	var err = 0
	var i int

	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			producto = &(*l)[i]
			break
		}
	}
	if i == len(*l) {
		err = 1
	}
	return producto, err
}

func (l *listaProductos) eliminarProducto(nombre string) {
	for i, producto := range *l {
		if producto.nombre == nombre {
			*l = append((*l)[:i], (*l)[i+1:]...)
			break
		}
	}
}

func (l *listaProductos) venderProducto(nombre string, cantidad int) {
	producto, err := l.buscarProducto(nombre)
	if err == 0 {
		if (*producto).cantidad >= cantidad {
			(*producto).cantidad -= cantidad
			if (*producto).cantidad == 0 {
				l.eliminarProducto(nombre)
				fmt.Printf("Producto agotado \n")
			} else {
				fmt.Printf("Venta \n")
			}
		} else {
			fmt.Printf(" No hay suficiente producto")
		}
	} else if err == 1 {
		fmt.Printf("Producto fuera del inventario \n", nombre)
	}
}

func (l *listaProductos) modificarPrecio(nombre string, nuevoPrecio int) {
	producto, err := l.buscarProducto(nombre)
	if err == 0 {
		(*producto).precio = nuevoPrecio
		fmt.Printf("Precio modificado\n")
	} else if err == 1 {
		fmt.Printf("Producto fuera del inventario")
	}
}

// Funcion para listar productos con existencia minima
func (l *listaProductos) listarProductosMinimos() listaProductos {
	productosMinimos := listaProductos{}

	for _, producto := range *l {
		if producto.cantidad < existenciaMinima {
			productosMinimos = append(productosMinimos, producto)
		}
	}
	return productosMinimos
}

// Funcion para aumentar inventario de productos con existencia minima
func (l *listaProductos) aumentarInventarioDeMinimos(listaMinimos listaProductos) {
	for _, prodMinimo := range listaMinimos {
		producto, err := l.buscarProducto(prodMinimo.nombre)
		//si se encuentra el producto
		if err == 0 {
			if (*producto).cantidad <= existenciaMinima {
				faltante := existenciaMinima - (*producto).cantidad
				fmt.Printf("Agregando %d unidades de '%s'\n", faltante, prodMinimo.nombre)
				(*producto).cantidad += faltante
			}
		}
	}
}

// Funcion para ordenar lista segun elemento y empezando de menor a mayor o al contrario
func ordenarListaPorElemento(lista listaProductos, elemento string, menor bool) {

	switch elemento {

	case "nombre":
		sort.SliceStable(lista, func(i, j int) bool {
			if menor {
				return lista[i].nombre < lista[j].nombre
			}
			return lista[i].nombre > lista[j].nombre
		})

	case "precio":
		sort.SliceStable(lista, func(i, j int) bool {
			if menor {
				return lista[i].precio < lista[j].precio
			}
			return lista[i].precio > lista[j].precio
		})
	case "cantidad":
		sort.SliceStable(lista, func(i, j int) bool {
			if menor {
				return lista[i].cantidad < lista[j].cantidad
			}
			return lista[i].cantidad > lista[j].cantidad
		})
	}
}

//FUNCIONES AUXILIARES

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
}

func llenarDatosInf() {
	lProductos.agregarProductosInfinita(
		producto{nombre: "arroz", cantidad: 20, precio: 5000},
		producto{nombre: "frijoles", cantidad: 15, precio: 6000},
		producto{nombre: "atun", cantidad: 5, precio: 5500},
		producto{nombre: "mantquilla", cantidad: 10, precio: 4500},
	)
}

func pruebaBuscar() {
	nombreProd := "arroz"
	prodEncontrado, err := lProductos.buscarProducto(nombreProd)
	if err == 0 {
		fmt.Printf("Producto encontrado: %s\n", prodEncontrado.nombre)
	} else if err == 1 {
		fmt.Printf("Producto no encontrado")
	}
}

func pruebaMinimos() {
	listaMinimos := lProductos.listarProductosMinimos()

	fmt.Println("Productos minimos:")
	for _, prod := range listaMinimos {
		fmt.Printf("Nombre: %s \n", prod.nombre)
	}
}

func main() {
	//Agregar datos
	llenarDatos()
	fmt.Println(lProductos)

	//Prueba buscar producto
	pruebaBuscar()

	//Prueba vender producto
	lProductos.venderProducto("arroz", 5)

	//Prueba modifixcar precio del producto
	lProductos.modificarPrecio("frijoles", 1800)

	//Prueba crear lista de minimos producto
	pruebaMinimos()

	//Prueba aumentar inventario de minimos
	listaMinimos := lProductos.listarProductosMinimos()
	lProductos.aumentarInventarioDeMinimos(listaMinimos)

	//Prueba ordenar lista
	fmt.Println("Lista inicial:")
	fmt.Println(lProductos)

	//Ordenar lista por nombre de manera ascendente
	fmt.Println("Lista ordenada:")
	ordenarListaPorElemento(lProductos, "nombre", true)
	fmt.Println(lProductos)
}
