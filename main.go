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

// Ejercicio 3: Copia independiente
//
// Se desarrolla la función copyArray, que crea un slice nuevo con make del
// mismo tamaño que la entrada y copia cada elemento mediante un ciclo, sin
// usar la función copy. Como make reserva un arreglo subyacente nuevo, la
// copia y el original quedan en memorias distintas y modificar uno no afecta
// al otro. En cambio, other := original no copia los elementos. Un slice
// guarda una referencia a su arreglo subyacente junto con su longitud y su
// capacidad, así que esa asignación solo copia la referencia y ambos slices
// terminan apuntando a los mismos datos. Por eso un cambio hecho con other
// también se ve en original. En Go no se llama a free, porque el recolector
// de basura recupera la memoria cuando el slice deja de ser alcanzable.

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
// Se desarrolla la función sortAscending, que ordena un slice de menor a
// mayor usando el algoritmo de inserción, sin usar sort ni slices.Sort. El
// algoritmo recorre el slice desde la segunda posición, guarda el elemento
// actual en una variable local y desplaza hacia la derecha los elementos
// anteriores que sean mayores, hasta encontrar el lugar donde insertarlo. La
// función no devuelve nada porque el descriptor del slice se pasa por valor,
// pero sus elementos pertenecen al arreglo subyacente compartido, así que los
// cambios se ven en el slice del llamador. Por eso se ordena una copia hecha
// con copyArray del ejercicio 3, para conservar la entrada original.

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
// Se desarrolla la función findPair, que busca en un slice ordenado de forma
// ascendente dos posiciones distintas cuyos valores sumen el objetivo. Usa
// dos índices, left en el inicio y right en el final, y en cada paso compara
// la suma de ambos con el objetivo. Si la suma es menor, se puede descartar
// el extremo izquierdo, porque al sumarlo con el mayor valor disponible ya no
// alcanza el objetivo y con cualquier otro valor la suma sería aún menor, así
// que left avanza. Si la suma es mayor, se descarta el extremo derecho por la
// razón contraria, porque ni con el menor valor disponible baja lo suficiente,
// así que right retrocede. El ciclo termina cuando encuentra el par o cuando
// los índices se cruzan, lo que da un recorrido O(n). La entrada no se
// modifica, solo se leen sus elementos, y si tiene menos de dos elementos el
// ciclo no se ejecuta y la función devuelve false.

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
// Se desarrolla la función threeSum, que imprime todas las ternas únicas de
// posiciones distintas cuya suma es cero y devuelve cuántas encontró. Si la
// entrada tiene menos de tres elementos devuelve cero. Para conservar la
// entrada se crea una copia con copyArray y se ordena con sortAscending. Luego
// se fija cada elemento sorted[i] como primer valor de la terna y se buscan
// los otros dos con la técnica de dos índices del ejercicio 5, con left
// después de i y right al final, buscando que sumen -sorted[i]. Al encontrar
// una terna no se detiene la búsqueda, porque con el mismo primer elemento
// puede haber otras ternas, así que mueve ambos índices y sigue. Las ternas
// duplicadas se evitan gracias al orden, ya que los valores iguales quedan
// juntos. Si sorted[i] es igual al anterior se salta, porque ese primer valor
// ya se procesó, y después de cada terna encontrada se saltan los valores
// repetidos en left y en right. El ciclo externo recorre n elementos y la
// búsqueda interna es O(n), así que la búsqueda posterior al ordenamiento es
// O(n²).

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
// Se desarrollan las funciones buildLeftMax y buildRightMax, que calculan
// para cada posición la mayor altura vista desde cada extremo, incluyendo la
// altura actual. Los slices de salida se crean antes con make, con la misma
// longitud que height, y se pasan como parámetros. Las funciones no necesitan
// devolverlos porque escriben en los elementos del arreglo subyacente que
// comparten con el llamador. buildLeftMax recorre de izquierda a derecha, y
// cada posición guarda el mayor entre la altura actual y el máximo de la
// posición anterior. buildRightMax hace lo mismo de derecha a izquierda,
// comparando con el máximo de la posición siguiente. Cada función hace un
// único recorrido, así que ambas son O(n). La entrada no se modifica y no
// puede estar vacía, porque cada recorrido empieza copiando la altura del
// extremo.

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
// Se desarrolla la función trappedWater, que calcula cuánta agua queda
// atrapada entre columnas de ancho uno. Crea leftMax y rightMax con make y
// los llena con buildLeftMax y buildRightMax del ejercicio 7. En cada
// posición, el agua depende del menor de los dos máximos porque el agua sube
// hasta que se derrama por el lado más bajo. La pared más alta de la
// izquierda y la más alta de la derecha forman un recipiente, y si el nivel
// pasara la menor de las dos, el agua se escaparía por ese lado hacia el
// extremo. Por eso el nivel en la posición es el menor de leftMax y rightMax,
// y el agua sobre la columna es ese nivel menos la altura actual. Como ambos
// máximos incluyen la altura actual, la resta nunca es negativa. El total se
// acumula en int64 porque con hasta 20 000 columnas y alturas de hasta
// 100 000 el resultado puede pasar el límite de un entero de 32 bits. La
// entrada no se modifica, el tiempo es O(n) porque hay tres recorridos
// separados y el espacio auxiliar es O(n) por los dos slices de máximos. La
// función printWaterDetail usa el mismo cálculo para mostrar el detalle por
// posición.

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
