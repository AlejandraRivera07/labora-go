package main

import "fmt"

type Planetas struct {
	planeta string
	lunas   int
}

func main() {

	planeta := []Planetas{
		{"Mercurio", 0},
		{"Venus", 0},
		{"Tierra", 1},
		{"Marte", 2},
		{"Júpiter", 63},
		{"Saturno", 62},
		{"Urano", 27},
		{"Neptuno", 13},
	}
	for i := 0; i < len(planeta); i++ {
		if planeta[i].lunas > 0 {
			fmt.Printf("%d. %s tiene %d lunas \n", i+1, planeta[i].planeta, planeta[i].lunas)
		} else {
			fmt.Printf("%d. %s no tiene lunas \n", i+1, planeta[i].planeta)

		}
	}

}
