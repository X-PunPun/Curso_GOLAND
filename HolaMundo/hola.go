package main

import "fmt"

func main() {
	fmt.Println("Hola Mundo")

	//variables
	var firstsNames, lastNames string
	var ages int

	firstsNames = "Andres"
	lastNames = "Marin"
	ages = 25

	var (
		firstsName = "Alex"
		lastName   = "Ruiz"
		Age        = 88
	)

	var (
		Primernombre, Segundonombre, edad = "camilo", "Rojas", 22
	)

	fmt.Println(firstsNames, lastNames, ages)
	fmt.Println(firstsName, lastName, Age)
	fmt.Println(Primernombre, Segundonombre, edad)
}
