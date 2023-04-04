/*package main
import "fmt"

func main() {
	var condicion bool
	condicion = true
	fmt.Printf("Data = %v, Type = %T", condicion, condicion)
}

*/

package main

import (
	"fmt"
)

func main() {

	var b1 bool = true // declaración tipada con valor inicializado
	var b2 = true      // declaración no-tipada con valor inicializado
	var b3 bool        // declaración tipada sin valor inicializado
	b4 := true         // declaración no-tipada con valor inicializado

	fmt.Println(b1) // Devuelve true
	fmt.Println(b2) // Devuelve true
	fmt.Println(b3) // Devuelve false
	fmt.Println(b4) // Devuelve true
}
