package main

import "fmt"

// Ejercicio 1: Intercambio de valores
//
// Se desarrolla la función swap, que intercambia los valores de dos
// variables. Recibe punteros porque Go pasa los argumentos por valor, así
// que si recibiera int trabajaría con copias y el cambio no se vería en main.
// Con punteros recibe las direcciones de las variables originales y puede
// modificarlas. Modificar a cambia solo la dirección guardada en el puntero
// local, mientras que modificar *a cambia el valor que está en esa dirección,
// que es la variable original de main. La variable auxiliar temp es local,
// por lo que solo existe dentro de swap y deja de ser accesible cuando la
// función termina.

// swap intercambia los valores a los que apuntan a y b.
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

// Ejercicio 2: Estadísticas de un arreglo
//
// Se desarrollan funciones que recorren un slice de enteros para calcular
// la suma, el mínimo, el máximo y la cantidad de números negativos, además
// de una función que imprime los elementos en orden inverso. Cada función
// recibe el slice como parámetro y usa variables locales para acumular su
// resultado, sin usar variables globales. El mínimo y el máximo empiezan con
// el primer elemento, por eso la entrada no puede estar vacía. El recorrido
// inverso solo avanza los índices del último al primero, así que la entrada
// no se modifica para imprimirla al revés.

// sumArray devuelve la suma de todos los elementos de values.
func sumArray(values []int) int {
	total := 0
	for i := 0; i < len(values); i++ {
		total += values[i]
	}
	return total
}

// minArray devuelve el menor elemento de values.
func minArray(values []int) int {
	minimum := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < minimum {
			minimum = values[i]
		}
	}
	return minimum
}

// maxArray devuelve el mayor elemento de values.
func maxArray(values []int) int {
	maximum := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > maximum {
			maximum = values[i]
		}
	}
	return maximum
}

// countNegatives devuelve cuántos elementos de values son menores que cero.
func countNegatives(values []int) int {
	count := 0
	for i := 0; i < len(values); i++ {
		if values[i] < 0 {
			count++
		}
	}
	return count
}

// printReverse imprime values del último al primer elemento.
func printReverse(values []int) {
	for i := len(values) - 1; i >= 0; i-- {
		fmt.Print(values[i])
		if i > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

// testStatistics ejecuta el caso de prueba del ejercicio 2.
func testStatistics() {
	fmt.Println("Ejercicio 2: Estadísticas de un arreglo")
	values := []int{4, -2, 0, 7, -5}
	fmt.Println("Entrada:", values)
	fmt.Println("Suma:", sumArray(values))
	fmt.Println("Mínimo:", minArray(values))
	fmt.Println("Máximo:", maxArray(values))
	fmt.Println("Negativos:", countNegatives(values))
	fmt.Print("Recorrido inverso: ")
	printReverse(values)
	fmt.Println("Entrada después del recorrido:", values)
	fmt.Println("Esperado: Suma: 4, Mínimo: -5, Máximo: 7, Negativos: 2")
	fmt.Println("Esperado: Recorrido inverso: -5 7 0 -2 4")
	fmt.Println()
}

func main() {
	testSwap()
	testStatistics()
}
