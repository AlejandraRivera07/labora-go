/*package main

import (
	"fmt"
)

func main() {

	var i, j, k = 1, 2, 3

	var (
		name       = "John Doe"
		occupation = "gardener"
	)

	fmt.Println(i, j, k)
	fmt.Printf("%s is a %s\n", name, occupation)
	fmt.Printf("%s tiene %d lunas")
}
*/

package main

import "fmt"

func main() {

	var age int = 34
	const WIDTH = 100

	age = 35
	age = 36

	// WIDTH = 101

	fmt.Println(age, WIDTH)
}
