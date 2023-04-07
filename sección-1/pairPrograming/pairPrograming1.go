package main

import (
	"fmt"
)

func main() {

	persona1 := Persona{nombre: "David", edad: 32, ciudad: "Bogota", telefono: 31398753}
	persona2 := Persona{nombre: "Ana", edad: 27, ciudad: "Medellin", telefono: 31938556}
	persona3 := Persona{nombre: "Miguel", edad: 21, ciudad: "Cali", telefono: 31789789}
	persona4 := Persona{nombre: "Camila", edad: 16, ciudad: "Bucaramanga", telefono: 318904985}
	persona5 := Persona{nombre: "Laura", edad: 28, ciudad: "Cartagena", telefono: 44567888}

	incrementarEdad(&persona1)
	fmt.Println("La nueva edad de la persona1  es: ", persona1.edad)

	personas := [5]Persona{persona1, persona2, persona3, persona4, persona5}

	buscar := buscarEdad(personas, 16)
	fmt.Printf("La persona %s tiene la edad ingresada  \n", buscar.nombre)

	personaNueva := crearPersona("Paula", 24, "Velez", 3456789)

	fmt.Println("La nueva persona fue creada \n", personaNueva)

	actualizarTelefono(&personaNueva, 33333333456)

	fmt.Printf("El numero de teléfono de %s se ha actualizado a %d \n", personaNueva.nombre, personaNueva.telefono)

}

type Persona struct {
	nombre   string
	edad     int
	ciudad   string
	telefono int64
}

func incrementarEdad(p *Persona) {
	p.edad++
}

func buscarEdad(personas [5]Persona, edad int) *Persona {
	for i := 0; i < len(personas); i++ {
		if personas[i].edad == edad {
			return &personas[i]
		}
	}

	return nil
}

func crearPersona(nombre string, edad int, ciudad string, telefono int64) Persona {
	nuevaPersona := Persona{nombre: nombre, edad: edad, ciudad: ciudad, telefono: telefono}
	return nuevaPersona
}

func actualizarTelefono(p *Persona, telefono int64) {
	p.telefono = telefono
}
