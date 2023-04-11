package main

import (
	"fmt"
	"sort"
)

type Estudiante struct {
	nombre string
	nota   float64
	codigo string
}

func main() {
	menu()

}

func menu() {
	// creo una lista donde cada elemento va a ser un estudiante con sus características
	var estudiantes []Estudiante
	numero := 0

	for numero != 8 {

		fmt.Println("___Menu___ ")
		fmt.Println("Por favor indicale al programa lo que quieres hacer")
		fmt.Println("1. Crear un estudiante")
		fmt.Println("2. Ordenar por nombre ")
		fmt.Println("3. Ordenar por nota ")
		fmt.Println("4. Ordenar  por código ")
		fmt.Println("5. Buscar  por nombre ")
		fmt.Println("6. Buscar  por nota ")
		fmt.Println("7. Buscar  por código ")
		fmt.Println("8. Salir ")
		fmt.Print("Ingrese una opción: ")
		fmt.Scan(&numero)

		switch numero {
		case 1:
			var estudiante Estudiante

			fmt.Println("Indica el nombre del estudiante ")
			fmt.Scan(&estudiante.nombre)
			fmt.Println("Indica la nota del estudiante ")
			fmt.Scan(&estudiante.nota)
			fmt.Println("Indica el código del estudiante ")
			fmt.Scan(&estudiante.codigo)

			estudiantes = append(estudiantes, estudiante)
			fmt.Println("Estudiante creado exitosamente.")

		case 2:
			sort.Slice(estudiantes, func(i int, j int) bool { return estudiantes[i].nombre < estudiantes[j].nombre })
			mostrarEstudiantes(estudiantes)
		case 3:
			sort.Slice(estudiantes, func(i int, j int) bool { return estudiantes[i].nota < estudiantes[j].nota })
			mostrarEstudiantes(estudiantes)
		case 4:
			sort.Slice(estudiantes, func(i int, j int) bool { return estudiantes[i].codigo < estudiantes[j].codigo })
			mostrarEstudiantes(estudiantes)
		case 5:
			var b_nombre string
			fmt.Println("Ingrese el nombre del estudiante")
			fmt.Scan(&b_nombre)

			buscar_nombre(estudiantes, b_nombre)

		case 6:
			var b_nota float64
			fmt.Println("Ingrese la nota del  estudiante")
			fmt.Scan(&b_nota)

			buscar_nota(estudiantes, b_nota)
		case 7:
			var b_codigo string
			fmt.Println("Ingrese el código del estudiante")
			fmt.Scan(&b_codigo)

			buscar_codigo(estudiantes, b_codigo)
		case 8:
			break
			fmt.Println("Adios!!!")
		}
	}
}

func mostrarEstudiantes(estudiantes []Estudiante) {
	fmt.Println("-- ESTUDIANTES ORDENADOS --")
	for _, estudiante := range estudiantes {
		fmt.Printf("Nombre: %s, Nota: %.2f, Código: %s\n", estudiante.nombre, estudiante.nota, estudiante.codigo)
		fmt.Printf("----------------------------------------------------------------------------------------------- \n")
		fmt.Print()
	}

}

func buscar_nombre(estudiantes []Estudiante, b_nombre string) {
	var encontrado bool
	for i := 0; i < len(estudiantes); i++ {
		if estudiantes[i].nombre == b_nombre {
			fmt.Printf("Se encontro el estudiante con Nombre: %s, nota: %.2f, codigo: %s\n", estudiantes[i].nombre, estudiantes[i].nota, estudiantes[i].codigo)
			encontrado = true
		}
	}
	if !encontrado {
		fmt.Printf("No existe ningun estudiante con nombre %s\n", b_nombre)
	}

}

func buscar_nota(estudiantes []Estudiante, b_nota float64) {
	var encontrado bool
	for i := 0; i < len(estudiantes); i++ {
		if estudiantes[i].nota == b_nota {
			fmt.Printf("Se encontro el estudiante con Nombre: %s, nota: %.2f, codigo: %s\n", estudiantes[i].nombre, estudiantes[i].nota, estudiantes[i].codigo)
			encontrado = true
		}
	}
	if !encontrado {
		fmt.Printf("No existe ningun estudiante con nota %s\n", b_nota)
	}

}

func buscar_codigo(estudiantes []Estudiante, b_codigo string) {
	var encontrado bool
	for i := 0; i < len(estudiantes); i++ {
		if estudiantes[i].codigo == b_codigo {
			fmt.Printf("Se encontro el estudiante con Nombre: %s, nota: %.2f, codigo: %s\n", estudiantes[i].nombre, estudiantes[i].nota, estudiantes[i].codigo)
			encontrado = true
		}
	}
	if !encontrado {
		fmt.Printf("No existe ningun estudiante con código  %s\n", b_codigo)
	}

}
