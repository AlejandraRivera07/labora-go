package main

import "fmt"

func main() {
	equipos := [5]int{1, 2, 3, 4, 5}
	titulos := [5]int{28, 14, 3, 7, 9}
	for i := 0; i < len(equipos); i++ {
		fmt.Printf("El equipo %v ha ganado %v títulos\n", equipos[i], titulos[i])
	}

}
