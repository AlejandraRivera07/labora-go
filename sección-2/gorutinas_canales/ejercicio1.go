package main

import (
	"fmt"
	"time"
)

func sumatoria(c chan<- int, start int, end int) {
	suma := 0
	for i := start; i < end; i++ {
		suma += i
	}
	c <- suma
	
}

func main() {

	c := make(chan int)
	

	
	go sumatoria(c, 1, 25)
	go sumatoria(c, 26, 50)
	go sumatoria(c, 51, 75)
	go sumatoria(c, 75, 100)

	time.Sleep(time.Second) 

	suma1 := <- c 
	suma2 := <- c
	suma3 := <- c
	suma4 := <- c

	fmt.Println("La suma es:", suma1+suma2+suma3+suma4)
}
