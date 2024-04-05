package main

import (
	"fmt"
)

func main() {
	arrayInitiation()
	arrayComp()
}

func arrayInitiation() {
	var x = [3]int{10, 20, 30}

	var y = [12]int{1, 5: 4, 6, 10: 100, 15}

	fmt.Println(x)
	fmt.Println(y)

	// Go has only one-dimensional arrays, but you can simulate multidimensional arrays:
	var z [2][3]int
	fmt.Println(z)

	// Like most languages, arrays in Go are read and written using bracket syntax:
	
	x[0] = 10
	fmt.Println(x[2])

}

func arrayComp() {
	// You can use == and != to compare two arrays. Arrays are equal if they are the same
	// length and contain equal values:
	var x = [...]int{1, 2, 3}
	var y = [3]int{1, 2, 3}
	fmt.Println(x == y) // prints true
}
