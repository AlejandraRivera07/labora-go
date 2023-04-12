package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	palabras := strings.Fields(s)
	conteo := make(map[string]int)
	for _, palabra := range palabras {
		conteo[palabra] += 1
	}

	return conteo
}

func main() {
	fmt.Print(WordCount("I ate a donut. Then I ate another donut."))
}
