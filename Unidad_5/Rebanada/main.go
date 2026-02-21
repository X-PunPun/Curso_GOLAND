package main

import "fmt"

func main() {
	diasSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles",
		"jueves", "Viernes", "Sabado"}

	fmt.Println(diasSemana)

	//cear una rebana a partir de otra
	diaRebanada := diasSemana[1:5]
	fmt.Println(diaRebanada)

	fmt.Println(diaRebanada)
	// append()

}
