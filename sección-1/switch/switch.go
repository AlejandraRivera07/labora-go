package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Ingresa un número del 1 al 7: ")
	fmt.Scan(&numero)

	switch numero {
	case 1:
		fmt.Println("1 corresponde con el día lunes")
	case 2:
		fmt.Println("2 corresponde con el día martes")
	case 3:
		fmt.Println("3 corresponde con el día miercoles")
	case 4:
		fmt.Println("4 corresponde con el día jueves")
	case 5:
		fmt.Println("5 corresponde con el día viernes")
	case 6:
		fmt.Println("6 corresponde con el día sabado")
	case 7:
		fmt.Println("7 corresponde con el día domingo")

	}

}
