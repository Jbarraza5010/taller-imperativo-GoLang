package main

import "fmt"

// Ejercicio 1: Intercambio de valores

// swap intercambia los valores de a y b.
//
// Recibe punteros porque Go pasa los argumentos por valor. Si recibiera
// int, trabajaría con copias y el cambio no se vería en main. Con punteros
// recibe las direcciones de las variables originales y puede modificarlas.
//
// Modificar a cambia solo la dirección guardada en el puntero local, por lo
// que la variable de main no cambia. Modificar *a cambia el valor que está
// en esa dirección, que es la variable original de main.
//
// temp es una variable local, así que solo existe dentro de swap y deja de
// ser accesible cuando la función termina.
func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// testSwap ejecuta el caso de prueba del ejercicio 1.
func testSwap() {
	fmt.Println("Ejercicio 1: Intercambio de valores")
	a := 8
	b := -3
	fmt.Printf("Antes:   a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("Después: a = %d, b = %d\n", a, b)
	fmt.Println("Esperado: a = -3, b = 8")
	fmt.Println()
}

func main() {
	testSwap()
}
