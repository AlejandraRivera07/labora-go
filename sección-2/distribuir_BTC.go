package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func contarVocales(nombre string) int {
	vocales := 0
	for _, letra := range strings.ToLower(nombre) {
		switch string(letra) {
		case "a":
			vocales += 1
		case "e":
			vocales += 1
		case "i":
			vocales += 2
		case "o":
			vocales += 3
		case "u":
			vocales += 4
		}
	}
	return vocales
}

func main() {
	for _, usuario := range users {
		asignarbtc := 0
		vocalesporusuario := contarVocales(usuario)

		//para garantizar que no hallan mas de 10 bts por usuairo
		if vocalesporusuario > 10 {
			asignarbtc = 10
		} else {
			asignarbtc = vocalesporusuario
		}
		distribution[usuario] = asignarbtc
		coins -= asignarbtc
	}

	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
}
