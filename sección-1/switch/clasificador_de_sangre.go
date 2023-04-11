package main

import (
	"fmt"
)

func clasificarSangre(tipoSangre string) string {
	var resultado string
	switch tipoSangre {
	case "AB+":
		fmt.Printf("Grupo sanguíneo AB, FACTOR Rh positivo \n")
	case "AB-":
		fmt.Printf("Grupo sanguíneo AB, FACTOR Rh negativo \n")
	case "A+":
		fmt.Printf("Grupo sanguíneo A, FACTOR Rh positivo \n")
	case "A-":
		fmt.Printf("Grupo sanguíneo A, FACTOR Rh negativo \n")
	case "B+":
		fmt.Printf("Grupo sanguíneo B, FACTOR Rh positivo \n")
	case "B-":
		fmt.Printf("Grupo sanguíneo B, FACTOR Rh negativo \n")
	case "O+":
		fmt.Printf("Grupo sanguíneo O, FACTOR Rh positivo \n")
	case "O-":
		fmt.Printf("Grupo sanguíneo O, FACTOR Rh negativo \n")

	}
	return resultado
}

func main() {
	var tipoSangre string
	fmt.Println("Escribe un tipo de sangre ")
	fmt.Scan(&tipoSangre)

	clasificarSangre(tipoSangre)
}
