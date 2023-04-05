package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Ingrese el primer número: ")
	fmt.Scanln(&num1)
	fmt.Print("Ingrese el segundo número: ")
	fmt.Scanln(&num2)

	sum, resta, multi, div := calcular(&num1, &num2)

	fmt.Println("Suma:", sum)
	fmt.Println("Resta:", resta)
	fmt.Println("Multiplicación:", multi)
	fmt.Println("División:", div)

}

func calcular(num1, num2 *int) (int, int, int, float64) {
	sum := *num1 + *num2
	resta := *num1 - *num2
	multi := *num1 * *num2
	div := float64(*num1) / float64(*num2)
	return sum, resta, multi, div
}
