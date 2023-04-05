package main

import "fmt"

func main() {
	var a int
	var b int

	a = 10
	b = 20
	var puntero *int = &a

	*puntero, b = b, *puntero

	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
