package main

import (
	"fmt"
)

func main() {

	peliculas := [10]string{"up", "coco", "enredados", "intensamente", "frozen", "encanto", "ratatouille", "cars", "soul", "la sirenita"}
	fmt.Println(peliculas)
	fmt.Printf("La segunda película del arreglo es: %s \n", peliculas[1])
	fmt.Printf("La longitud del arreglo es: %v \n", len(peliculas))
}
