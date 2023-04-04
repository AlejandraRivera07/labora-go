package main

import "fmt"

func main() {
	restaurantes := [3]int{1, 2, 3}
	calificaciones := [3]float32{4.5, 3.8, 2.6}
	for i := 0; i < len(restaurantes); i++ {
		fmt.Printf("El restaurante %v tiene una calificación de %v \n", restaurantes[i], calificaciones[i])
	}

}
