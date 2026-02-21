package main

import "fmt"

func main() {
	fmt.Println("\nllamada de matrices")
	var a [5]int //martris con 5 elementos int
	//modificar datos de matriz
	a[0] = 10
	a[1] = 20
	a[2] = 30
	a[3] = 40
	a[4] = 50
	fmt.Println(a)

	//crear matriz con datos existentes.
	var as = [5]int{10, 20, 30, 40, 50}
	fmt.Println(as)

	//crear matiz con datos indifinidos
	var asd = [...]int{10, 20, 30, 40, 50, 80, 100}
	fmt.Println(asd)

	//ejemplificacion de como adceder a un elemento por index
	fmt.Print("\nMatrices llamadas por idices 3 0 6")
	fmt.Println(a[3])
	fmt.Println(as[0])
	fmt.Println(asd[6])

	//Ejemplificion con for
	fmt.Println("\nEjemplificacion de cadena con For")
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	//Ejemplificacion de indice = valor
	fmt.Println("\nEjemplificacion Indice = valor")
	for index, value := range a {
		fmt.Printf("indice: %d, Valor %d \n", index, value)
	}

	var matriz = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Printf("\nMatriz bidimencional:  %d", matriz)
}
