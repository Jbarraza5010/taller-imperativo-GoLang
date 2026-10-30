# taller-imperativo-GoLang

Taller que aplica el paradigma imperativo mediante funciones, estructuras de control, arreglos, datos compuestos y manejo de memoria, hasta resolver los problemas 3Sum y Trapping Rain Water.

Instituto Tecnológico de Costa Rica, Ingeniería en Computadores, curso de Paradigmas de Programación.

## Integrantes

| Nombre | Carnet |
| --- | --- |
| Javier Alejandro Barraza Arias | 2022114377 |

## Lenguaje

Go. Todo el código está en `main.go` y no usa bibliotecas externas, solo el paquete estándar `fmt`.

## Instrucciones de ejecución

1. Instalar Go, por ejemplo en Ubuntu con `sudo snap install go --classic`, y comprobar la instalación con `go version`.
2. Desde la carpeta del repositorio, ejecutar:

```bash
go run main.go
```

El programa corre los casos de prueba de los ocho ejercicios en orden. Cada caso imprime el resultado obtenido y, debajo, el resultado esperado según el enunciado. No se pide entrada al usuario, porque las entradas están definidas en el código.

## Ejercicios

| Etapa | Ejercicio | Funciones principales |
| --- | --- | --- |
| 1 | Intercambio de valores | `swap` |
| 2 | Estadísticas de un arreglo | `sumArray`, `minArray`, `maxArray`, `countNegatives`, `printReverse` |
| 3 | Copia independiente | `copyArray` |
| 4 | Ordenar una copia | `sortAscending` (inserción) |
| 5 | Dos valores que suman un objetivo | `findPair` |
| 6 | 3Sum | `threeSum` |
| 7 | Máximos desde ambos extremos | `buildLeftMax`, `buildRightMax` |
| 8 | Trapping Rain Water | `trappedWater`, `printWaterDetail` |

Cada ejercicio tiene en el código un comentario que explica lo que se desarrolla y una función de prueba que lo ejecuta desde `main`.

## Resultados de las pruebas

Todos los casos coinciden con los resultados esperados del enunciado.

### Fundamentos y arreglos

| Ejercicio | Entrada | Resultado obtenido |
| --- | --- | --- |
| 1 | a = 8, b = -3 | a = -3, b = 8 |
| 2 | [4, -2, 0, 7, -5] | Suma 4, mínimo -5, máximo 7, negativos 2, inverso -5 7 0 -2 4 |
| 3 | [4, -2, 0, 7, -5] | Original [4, -2, 0, 7, -5], copia [99, -2, 0, 7, -5] |
| 4 | [4, -2, 0, 7, -5] | Ordenado [-5, -2, 0, 4, 7], original sin cambios |

### Búsqueda de pares

| Arreglo | Objetivo | Resultado obtenido |
| --- | --- | --- |
| [-4, -1, 0, 2, 5] | 1 | (-4, 5) |
| [-4, -1, 0, 2, 5] | 8 | No existe |
| [0, 0] | 0 | (0, 0) |
| [3] | 6 | No existe, por tamaño insuficiente |

### 3Sum

| Entrada | Ternas obtenidas | Cantidad |
| --- | --- | --- |
| [-1, 0, 1, 2, -1, -4] | (-1, -1, 2) y (-1, 0, 1) | 2 |
| [0, 1, 1] | Ninguna | 0 |
| [0, 0, 0] | (0, 0, 0) | 1 |
| [0, 0, 0, 0] | (0, 0, 0) | 1 |
| [-2, 0, 0, 2, 2] | (-2, 0, 2) | 1 |

En todos los casos la entrada se conserva después de la búsqueda.

### Máximos acumulados

```
height:   [3 0 2 0 4]
leftMax:  [3 3 3 3 4]
rightMax: [4 4 4 4 4]
```

### Agua atrapada

| Alturas | Agua obtenida |
| --- | --- |
| [5] | 0 |
| [0, 1, 2, 3] | 0 |
| [3, 2, 1, 0] | 0 |
| [2, 2, 2] | 0 |
| [3, 0, 3] | 3 |
| [3, 0, 2, 0, 4] | 7 |
| [4, 2, 0, 3, 2, 5] | 9 |
| [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1] | 6 |
| [5, 0, 0, 2] (caso propio) | 4 |

Detalle por posición para [3, 0, 2, 0, 4]:

| Posición | Altura | leftMax | rightMax | Menor máximo | Agua |
| --- | --- | --- | --- | --- | --- |
| 0 | 3 | 3 | 4 | 3 | 0 |
| 1 | 0 | 3 | 4 | 3 | 3 |
| 2 | 2 | 3 | 4 | 3 | 1 |
| 3 | 0 | 3 | 4 | 3 | 3 |
| 4 | 4 | 4 | 4 | 4 | 0 |
| | | | | Total | 7 |

**Caso propio [5, 0, 0, 2].** La pared izquierda mide 5 y la derecha solo 2. En las posiciones 1 y 2 el menor de los máximos es 2, así que el agua sube hasta 2 y no hasta 5, porque por encima de esa altura se derramaría por la derecha. Cada posición guarda 2 unidades y el total esperado es 4, que coincide con el obtenido. Se eligió este caso porque los dos máximos son muy distintos y muestra con claridad que el nivel lo decide el lado más bajo.

<details>
<summary>Salida completa de <code>go run main.go</code></summary>

```
Ejercicio 1: Intercambio de valores
Antes:   a = 8, b = -3
Después: a = -3, b = 8
Esperado: a = -3, b = 8

Ejercicio 2: Estadísticas de un arreglo
Entrada: [4 -2 0 7 -5]
Suma: 4
Mínimo: -5
Máximo: 7
Negativos: 2
Recorrido inverso: -5 7 0 -2 4
Entrada después del recorrido: [4 -2 0 7 -5]
Esperado: Suma: 4, Mínimo: -5, Máximo: 7, Negativos: 2
Esperado: Recorrido inverso: -5 7 0 -2 4

Ejercicio 3: Copia independiente
Original: [4 -2 0 7 -5]
Copia:    [99 -2 0 7 -5]
Esperado: Original [4 -2 0 7 -5], Copia [99 -2 0 7 -5]
Comparación con other := original
Original: [99 -2 0 7 -5]
Other:    [99 -2 0 7 -5]
El original también cambia porque ambos comparten los datos

Ejercicio 4: Ordenar una copia
Original: [4 -2 0 7 -5]
Ordenado: [-5 -2 0 4 7]
Esperado: Ordenado [-5 -2 0 4 7]

Ejercicio 5: Dos valores que suman un objetivo
Arreglo: [-4 -1 0 2 5] Objetivo: 1
Par encontrado: (-4, 5)
Esperado: (-4, 5)
Arreglo: [-4 -1 0 2 5] Objetivo: 8
No existe un par
Esperado: No existe
Arreglo: [0 0] Objetivo: 0
Par encontrado: (0, 0)
Esperado: (0, 0)
Arreglo: [3] Objetivo: 6
No existe un par
Esperado: No existe

Ejercicio 6: 3Sum
Entrada: [-1 0 1 2 -1 -4]
(-1, -1, 2)
(-1, 0, 1)
Cantidad de ternas: 2
Entrada después de la búsqueda: [-1 0 1 2 -1 -4]
Esperado: (-1, -1, 2) y (-1, 0, 1)
Entrada: [0 1 1]
Cantidad de ternas: 0
Entrada después de la búsqueda: [0 1 1]
Esperado: Ninguna
Entrada: [0 0 0]
(0, 0, 0)
Cantidad de ternas: 1
Entrada después de la búsqueda: [0 0 0]
Esperado: (0, 0, 0)
Entrada: [0 0 0 0]
(0, 0, 0)
Cantidad de ternas: 1
Entrada después de la búsqueda: [0 0 0 0]
Esperado: Una única terna (0, 0, 0)
Entrada: [-2 0 0 2 2]
(-2, 0, 2)
Cantidad de ternas: 1
Entrada después de la búsqueda: [-2 0 0 2 2]
Esperado: Una única terna (-2, 0, 2)

Ejercicio 7: Máximos desde ambos extremos
height:   [3 0 2 0 4]
leftMax:  [3 3 3 3 4]
rightMax: [4 4 4 4 4]
Esperado: leftMax [3 3 3 3 4], rightMax [4 4 4 4 4]

Ejercicio 8: Trapping Rain Water
Alturas: [5] Agua total: 0 Esperado: 0
Alturas: [0 1 2 3] Agua total: 0 Esperado: 0
Alturas: [3 2 1 0] Agua total: 0 Esperado: 0
Alturas: [2 2 2] Agua total: 0 Esperado: 0
Alturas: [3 0 3] Agua total: 3 Esperado: 3
Alturas: [3 0 2 0 4] Agua total: 7 Esperado: 7
Alturas: [4 2 0 3 2 5] Agua total: 9 Esperado: 9
Alturas: [0 1 0 2 1 0 1 3 2 1 2 1] Agua total: 6 Esperado: 6
Detalle por posición para [3, 0, 2, 0, 4]
Posición 0: altura 3, leftMax 3, rightMax 4, agua 0
Posición 1: altura 0, leftMax 3, rightMax 4, agua 3
Posición 2: altura 2, leftMax 3, rightMax 4, agua 1
Posición 3: altura 0, leftMax 3, rightMax 4, agua 3
Posición 4: altura 4, leftMax 4, rightMax 4, agua 0
Entrada después del cálculo: [3 0 2 0 4]
Caso propio
Alturas: [5 0 0 2] Agua total: 4 Esperado: 4
Posición 0: altura 5, leftMax 5, rightMax 5, agua 0
Posición 1: altura 0, leftMax 5, rightMax 2, agua 2
Posición 2: altura 0, leftMax 5, rightMax 2, agua 2
Posición 3: altura 2, leftMax 5, rightMax 2, agua 0
```

</details>

## Cómo se evitan las ternas duplicadas en 3Sum

La función `threeSum` primero crea una copia de la entrada con `copyArray` y la ordena con `sortAscending`. Al quedar ordenada, los valores iguales quedan uno junto al otro, y eso permite reconocer y saltar los repetidos.

La búsqueda fija cada elemento `sorted[i]` como primer valor de la terna y busca los otros dos con dos índices, `left` justo después de `i` y `right` al final. Los duplicados se evitan en tres puntos:

1. **En el primer valor.** Si `sorted[i]` es igual a `sorted[i-1]`, se salta esa posición, porque todas las ternas que empiezan con ese valor ya se encontraron en la iteración anterior.
2. **En el segundo valor.** Después de encontrar una terna, `left` avanza mientras el nuevo valor sea igual al anterior.
3. **En el tercer valor.** De la misma forma, `right` retrocede mientras el nuevo valor sea igual al anterior.

Al encontrar una terna la búsqueda no se detiene, porque con el mismo primer valor puede haber otras ternas distintas. Por ejemplo, con `[-1, 0, 1, 2, -1, -4]` ordenado como `[-4, -1, -1, 0, 1, 2]`, el primer `-1` produce `(-1, -1, 2)` y `(-1, 0, 1)`, y el segundo `-1` se salta. Así cada terna aparece una sola vez. El ciclo externo recorre n posiciones y la búsqueda con dos índices es O(n), así que la búsqueda posterior al ordenamiento es O(n²).

## Por qué el agua depende del menor de los dos máximos

En cada posición, la columna más alta desde el inicio hasta ahí (`leftMax`) y la más alta desde ahí hasta el final (`rightMax`) funcionan como las paredes de un recipiente. Si el agua subiera por encima del menor de los dos máximos, en ese lado no habría ninguna columna lo bastante alta para contenerla, así que se derramaría hasta salir por el extremo. Y no puede quedar por debajo de ese nivel, porque a ambos lados hay una pared al menos así de alta que la retiene. Por eso el nivel del agua en la posición es exactamente el menor de los dos máximos, y no el mayor.

El agua sobre la columna es ese nivel menos la altura de la columna:

```
agua[i] = menor(leftMax[i], rightMax[i]) - height[i]
```

Como ambos máximos incluyen la altura actual, el resultado nunca es negativo. Si la columna es la más alta de uno de sus lados, el nivel coincide con su altura y no guarda agua. El caso propio `[5, 0, 0, 2]` lo muestra: aunque a la izquierda hay una pared de 5, el agua solo sube hasta 2, que es la altura de la pared derecha.

El total se acumula en `int64`, como pide el enunciado. Con hasta 20 000 columnas y alturas de hasta 100 000, el peor caso es de 19 998 × 100 000 = 1 999 800 000 unidades, muy cerca del límite de un entero de 32 bits (2 147 483 647). Usar `int64` deja un margen amplio y garantiza que la suma no se desborde, sin depender del tamaño de `int` en la plataforma.

## Ejemplo de alcance de variables

Tomado de la función `threeSum`:

```go
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
			...
		}
	}
	return count
}
```

- `values` es un parámetro y `sorted` y `count` son variables locales de la función. Existen durante toda la llamada a `threeSum` y no se pueden usar desde otras funciones.
- `i` se declara en el encabezado del `for`, así que solo existe dentro de ese ciclo.
- `left` y `right` se declaran dentro del cuerpo del ciclo externo. Se crean de nuevo en cada iteración con los valores iniciales de esa iteración, y no existen fuera del ciclo.
- `sum` se declara dentro del ciclo interno, así que se recalcula en cada paso y no se puede usar fuera de él.
- `count` se declara fuera de los ciclos porque tiene que acumular el total de todas las iteraciones y seguir existiendo para el `return`.

El programa no usa variables globales. Todos los datos se pasan mediante parámetros, como pide el enunciado.

## Manejo de memoria en Go y en C

| Aspecto | C | Go |
| --- | --- | --- |
| Reservar memoria | `malloc(size * sizeof(int))` | `make([]int, size)` |
| Valor inicial | Indefinido | Cero en todos los elementos |
| Fallo de reserva | `malloc` devuelve `NULL` y hay que comprobarlo | No se comprueba, el taller no lo exige |
| Liberar memoria | Manual con `free` | Automático con el recolector de basura |
| Tamaño del arreglo | Se pasa como parámetro aparte | Va dentro del slice y se obtiene con `len` |
| Arreglo en una función | Se pasa un puntero al primer elemento | Se pasa el descriptor del slice por valor, que comparte el arreglo subyacente |

En C el programador es responsable de toda la memoria dinámica. Cada `malloc` debe comprobarse y cada bloque reservado debe liberarse con `free`, y un error en eso produce fugas de memoria o accesos a memoria ya liberada. En Go, `make` reserva el arreglo subyacente e inicializa los elementos en cero, y el recolector de basura recupera la memoria cuando ningún slice la alcanza, así que no existe `free`.

Lo que sí se mantiene en ambos lenguajes es la idea de memoria compartida. En C, copiar un puntero no copia los datos, y en Go `other := original` tampoco, porque solo copia la referencia al arreglo subyacente. Por eso en el ejercicio 3 la copia independiente se hace con `make` y un ciclo, igual que en C se haría con `malloc` y un ciclo.
