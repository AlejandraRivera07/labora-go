package main

import "fmt"

func main() {
	planets := [8]string{"Mercurio", "Venus", "Tierra", "Marte", "Júpiter", "Saturno", "Urano", "Neptuno"}
	lunas := [8]int{0, 0, 1, 2, 63, 62, 27, 13}
	for i := 0; i < len(planets); i++ {
		fmt.Printf("%s tiene %v lunas\n", planets[i], lunas[i])
	}

}
