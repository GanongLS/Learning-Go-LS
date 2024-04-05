package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ifConstruct()

}

func ifConstruct() {
	fmt.Println("If Construct")
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	fmt.Println("n:", n)

	if m := rand.Intn(10); m == 0 { // Scoping a variable to an if statement
		fmt.Println("That's too low")
	} else if m > 5 {
		fmt.Println("That's too big:", m)
	} else {
		fmt.Println("That's a good number:", m)
	}
	// Having this special scope is handy. It lets you create variables that are available only where they are needed. Once the series of if/else statements ends, n is undefined.

	// fmt.Println("n:", m) // undefined m

	fmt.Println("")

}
