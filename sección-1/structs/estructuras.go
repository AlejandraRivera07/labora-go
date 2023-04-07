package main

import (
	"fmt"
)

func main() {
	persona1 := Persona{"David", 32, "Bogota", 31398753}
	persona2 := Persona{"Ana", 27, "Medellin", 31938556}

	fmt.Println("Persona 1", persona1)
	fmt.Println("Persona 2", persona2)

	cambiarCiudad(&persona1, "Cali")

	fmt.Printf("persona1 con ciudad actualizada: %v\n", persona1)
	fmt.Println("Persona2 ", persona2)
}

type Persona struct {
	Nombre   string
	Edad     int
	Ciudad   string
	Telefono int
}

func cambiarCiudad(persona *Persona, ciudad string) {
	if persona.Ciudad != ciudad {
		persona.Ciudad = ciudad
	}

}
