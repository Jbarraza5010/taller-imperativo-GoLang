package main

import "fmt"

// Ejercicio 1: Intercambio de valores
//
// La función swap intercambia dos valores. Recibe punteros porque Go pasa
// copias de los argumentos, y así puede modificar las variables originales.
// Cambiar a solo cambia el puntero local, mientras que cambiar *a modifica la
// variable de main. La variable temp es local y solo existe dentro de swap.

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
// Funciones que recorren un slice para calcular la suma, el mínimo, el
// máximo y los negativos, y otra que lo imprime al revés. Cada una recibe el
// slice como parámetro y usa variables locales. El recorrido inverso solo
// mueve los índices, así que la entrada no se modifica.

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

// Ejercicio 3: Copia independiente
//
// La función copyArray crea un slice nuevo con make y copia los elementos
// con un ciclo, sin usar copy. Como make reserva memoria nueva, cambiar la
// copia no afecta al original. En cambio, other := original solo copia la
// referencia al arreglo subyacente, así que ambos comparten los mismos datos.
// En Go no se llama a free porque el recolector libera la memoria.

// copyArray devuelve un slice nuevo con los mismos elementos de values.
func copyArray(values []int) []int {
	result := make([]int, len(values))
	for i := 0; i < len(values); i++ {
		result[i] = values[i]
	}
	return result
}

// testCopy ejecuta el caso de prueba del ejercicio 3.
func testCopy() {
	fmt.Println("Ejercicio 3: Copia independiente")
	original := []int{4, -2, 0, 7, -5}
	independent := copyArray(original)
	independent[0] = 99
	fmt.Println("Original:", original)
	fmt.Println("Copia:   ", independent)
	fmt.Println("Esperado: Original [4 -2 0 7 -5], Copia [99 -2 0 7 -5]")

	fmt.Println("Comparación con other := original")
	shared := []int{4, -2, 0, 7, -5}
	other := shared
	other[0] = 99
	fmt.Println("Original:", shared)
	fmt.Println("Other:   ", other)
	fmt.Println("El original también cambia porque ambos comparten los datos")
	fmt.Println()
}

// Ejercicio 4: Ordenar una copia
//
// La función sortAscending ordena de menor a mayor con inserción, sin usar
// sort ni slices.Sort. Cada elemento se inserta en su lugar desplazando a la
// derecha los mayores que tiene antes. No devuelve nada porque modifica el
// arreglo subyacente que comparte con el llamador, por eso se ordena una
// copia hecha con copyArray y la entrada se conserva.

// sortAscending ordena values de menor a mayor mediante inserción.
func sortAscending(values []int) {
	for i := 1; i < len(values); i++ {
		current := values[i]
		j := i - 1
		for j >= 0 && values[j] > current {
			values[j+1] = values[j]
			j--
		}
		values[j+1] = current
	}
}

// testSort ejecuta el caso de prueba del ejercicio 4.
func testSort() {
	fmt.Println("Ejercicio 4: Ordenar una copia")
	original := []int{4, -2, 0, 7, -5}
	sorted := copyArray(original)
	sortAscending(sorted)
	fmt.Println("Original:", original)
	fmt.Println("Ordenado:", sorted)
	fmt.Println("Esperado: Ordenado [-5 -2 0 4 7]")
	fmt.Println()
}

// Ejercicio 5: Dos valores que suman un objetivo
//
// La función findPair busca en un slice ordenado dos posiciones que sumen el
// objetivo, con un índice en cada extremo. Si la suma es menor, el extremo
// izquierdo se descarta porque ni con el mayor valor alcanza, y left avanza.
// Si es mayor, el derecho se descarta porque ni con el menor baja, y right
// retrocede. Es O(n), no modifica la entrada y con menos de dos elementos
// devuelve false.

// findPair imprime un par de values que suma target y devuelve si existe.
func findPair(values []int, target int) bool {
	left := 0
	right := len(values) - 1
	for left < right {
		sum := values[left] + values[right]
		if sum == target {
			fmt.Printf("Par encontrado: (%d, %d)\n", values[left], values[right])
			return true
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return false
}

// runPairTest ejecuta un caso de findPair y muestra el resultado esperado.
func runPairTest(values []int, target int, expected string) {
	fmt.Println("Arreglo:", values, "Objetivo:", target)
	found := findPair(values, target)
	if !found {
		fmt.Println("No existe un par")
	}
	fmt.Println("Esperado:", expected)
}

// testFindPair ejecuta las cuatro pruebas de búsqueda de pares.
func testFindPair() {
	fmt.Println("Ejercicio 5: Dos valores que suman un objetivo")
	runPairTest([]int{-4, -1, 0, 2, 5}, 1, "(-4, 5)")
	runPairTest([]int{-4, -1, 0, 2, 5}, 8, "No existe")
	runPairTest([]int{0, 0}, 0, "(0, 0)")
	runPairTest([]int{3}, 6, "No existe")
	fmt.Println()
}

// Ejercicio 6: 3Sum
//
// La función threeSum imprime las ternas únicas que suman cero y devuelve su
// cantidad, o cero si hay menos de tres elementos. Ordena una copia para
// conservar la entrada, fija cada sorted[i] y busca los otros dos con dos
// índices como en el ejercicio 5. Al encontrar una terna sigue buscando,
// porque puede haber más con el mismo primer valor. Los duplicados se evitan
// saltando los valores repetidos en i, left y right, que quedan juntos por el
// orden. La búsqueda es O(n²).

// threeSum imprime las ternas únicas que suman cero y devuelve su cantidad.
func threeSum(values []int) int {
	if len(values) < 3 {
		return 0
	}
	sorted := copyArray(values)
	sortAscending(sorted)
	count := 0
	for i := 0; i < len(sorted)-2; i++ {
		if i > 0 && sorted[i] == sorted[i-1] {
			continue
		}
		left := i + 1
		right := len(sorted) - 1
		for left < right {
			sum := sorted[i] + sorted[left] + sorted[right]
			if sum == 0 {
				fmt.Printf("(%d, %d, %d)\n", sorted[i], sorted[left], sorted[right])
				count++
				left++
				right--
				for left < right && sorted[left] == sorted[left-1] {
					left++
				}
				for left < right && sorted[right] == sorted[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return count
}

// runThreeSumTest ejecuta un caso de threeSum y muestra el resultado esperado.
func runThreeSumTest(values []int, expected string) {
	fmt.Println("Entrada:", values)
	count := threeSum(values)
	fmt.Println("Cantidad de ternas:", count)
	fmt.Println("Entrada después de la búsqueda:", values)
	fmt.Println("Esperado:", expected)
}

// testThreeSum ejecuta los cinco casos de prueba de 3Sum.
func testThreeSum() {
	fmt.Println("Ejercicio 6: 3Sum")
	runThreeSumTest([]int{-1, 0, 1, 2, -1, -4}, "(-1, -1, 2) y (-1, 0, 1)")
	runThreeSumTest([]int{0, 1, 1}, "Ninguna")
	runThreeSumTest([]int{0, 0, 0}, "(0, 0, 0)")
	runThreeSumTest([]int{0, 0, 0, 0}, "Una única terna (0, 0, 0)")
	runThreeSumTest([]int{-2, 0, 0, 2, 2}, "Una única terna (-2, 0, 2)")
	fmt.Println()
}

// Ejercicio 7: Máximos desde ambos extremos
//
// Las funciones buildLeftMax y buildRightMax guardan en cada posición la
// mayor altura vista desde cada extremo, incluyendo la actual. Los slices de
// salida se crean antes con make y las funciones escriben en ellos, porque
// comparten el arreglo subyacente con el llamador. Cada una compara la altura
// actual con el máximo de la posición vecina y hace un solo recorrido O(n).

// buildLeftMax llena leftMax con el máximo de height desde el inicio.
func buildLeftMax(height []int, leftMax []int) {
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		if height[i] > leftMax[i-1] {
			leftMax[i] = height[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}
}

// buildRightMax llena rightMax con el máximo de height desde el final.
func buildRightMax(height []int, rightMax []int) {
	last := len(height) - 1
	rightMax[last] = height[last]
	for i := last - 1; i >= 0; i-- {
		if height[i] > rightMax[i+1] {
			rightMax[i] = height[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}
}

// testMaxArrays ejecuta el caso de prueba del ejercicio 7.
func testMaxArrays() {
	fmt.Println("Ejercicio 7: Máximos desde ambos extremos")
	height := []int{3, 0, 2, 0, 4}
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	buildLeftMax(height, leftMax)
	buildRightMax(height, rightMax)
	fmt.Println("height:  ", height)
	fmt.Println("leftMax: ", leftMax)
	fmt.Println("rightMax:", rightMax)
	fmt.Println("Esperado: leftMax [3 3 3 3 4], rightMax [4 4 4 4 4]")
	fmt.Println()
}

// Ejercicio 8: Trapping Rain Water
//
// La función trappedWater calcula el agua atrapada entre las columnas usando
// leftMax y rightMax del ejercicio 7. El agua de cada posición depende del
// menor de los dos máximos porque, si subiera más, se derramaría por el lado
// más bajo. Esa altura menos la columna actual da el agua de la posición. El
// total se acumula en int64 porque el peor caso queda cerca del límite de 32
// bits. Usa tiempo O(n), espacio auxiliar O(n) y no modifica la entrada.

// waterAt devuelve el agua sobre la posición i a partir de los dos máximos.
func waterAt(height []int, leftMax []int, rightMax []int, i int) int {
	level := leftMax[i]
	if rightMax[i] < level {
		level = rightMax[i]
	}
	return level - height[i]
}

// trappedWater devuelve el total de agua atrapada sobre height.
func trappedWater(height []int) int64 {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	buildLeftMax(height, leftMax)
	buildRightMax(height, rightMax)
	total := int64(0)
	for i := 0; i < len(height); i++ {
		amount := waterAt(height, leftMax, rightMax, i)
		total += int64(amount)
	}
	return total
}

// printWaterDetail imprime el cálculo de agua para cada posición de height.
func printWaterDetail(height []int) {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	buildLeftMax(height, leftMax)
	buildRightMax(height, rightMax)
	for i := 0; i < len(height); i++ {
		amount := waterAt(height, leftMax, rightMax, i)
		fmt.Printf("Posición %d: altura %d, leftMax %d, rightMax %d, agua %d\n",
			i, height[i], leftMax[i], rightMax[i], amount)
	}
}

// runWaterTest ejecuta un caso de trappedWater y muestra el resultado esperado.
func runWaterTest(height []int, expected int64) {
	total := trappedWater(height)
	fmt.Println("Alturas:", height, "Agua total:", total, "Esperado:", expected)
}

// testTrappedWater ejecuta los ocho casos de agua atrapada y un caso propio.
func testTrappedWater() {
	fmt.Println("Ejercicio 8: Trapping Rain Water")
	runWaterTest([]int{5}, 0)
	runWaterTest([]int{0, 1, 2, 3}, 0)
	runWaterTest([]int{3, 2, 1, 0}, 0)
	runWaterTest([]int{2, 2, 2}, 0)
	runWaterTest([]int{3, 0, 3}, 3)
	runWaterTest([]int{3, 0, 2, 0, 4}, 7)
	runWaterTest([]int{4, 2, 0, 3, 2, 5}, 9)
	runWaterTest([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6)

	fmt.Println("Detalle por posición para [3, 0, 2, 0, 4]")
	detail := []int{3, 0, 2, 0, 4}
	printWaterDetail(detail)
	fmt.Println("Entrada después del cálculo:", detail)

	// Caso propio: la pared izquierda mide 5 pero la derecha solo 2, así que
	// el agua en las posiciones 1 y 2 sube hasta 2 y no hasta 5. Cada una
	// guarda 2 unidades y el total esperado es 4.
	fmt.Println("Caso propio")
	runWaterTest([]int{5, 0, 0, 2}, 4)
	printWaterDetail([]int{5, 0, 0, 2})
	fmt.Println()
}

func main() {
	testSwap()
	testStatistics()
	testCopy()
	testSort()
	testFindPair()
	testThreeSum()
	testMaxArrays()
	testTrappedWater()
}
