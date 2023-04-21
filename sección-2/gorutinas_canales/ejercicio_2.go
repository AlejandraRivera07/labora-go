package main

import (
	"fmt"
)

func multiplication(c chan<- [][]int, a [][]int, b [][]int) {
	rowsA := len(a)
	colsA := len(a[0])
	colsB := len(b[0])

	resultado := make([][]int, rowsA)

	for i := range resultado {
		resultado[i] = make([]int, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				resultado[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	c <- resultado
}

func main() {

	a := [][]int{{1, 2}, {4, 5}}
	b := [][]int{{7, 8}, {4, 7}}

	c := make(chan [][]int)

	go multiplication(c, a, b)
	result := <-c
	fmt.Println("El resultado de la multiplicación es: \n")
	for _, row := range result {
		fmt.Println(row)
	}

}
