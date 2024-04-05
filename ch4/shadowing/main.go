package main

import (
	"fmt"
)

func main() {
	shadowing()
	anotherShadowing()
	messingWithUniverseBlock()
}

func messingWithUniverseBlock() {
	fmt.Println(true)
	true := 10
	fmt.Println(true)

	// You must be very careful to never redefine any identifiers in the universe block.
}

// func unintendedShadowing() {
// 	x := 10
// 	fmt.Println(x)
// 	fmt := "oops"
// 	// fmt.Println(fmt) ga bisa jalan, fmt udah ke shadow.
// }

func anotherShadowing() {
	x := 10
	if x > 5 {
		x, y := 5, 20 // notasi := cukup tricky karena bisa digunakan cukup asal ada satu variable baru, sehingga bisa membuat hidden shadowing, unintended shadowing
		fmt.Println(x, y)
	}
	fmt.Println(x)
}

func shadowing() {
	fmt.Println("Shadowing")
	// when you have a declaration with the same name as an identifier in a containing block? If you do that, you shadow the identifier created in the outer block.

	x := 10
	if x > 5 {
		fmt.Println(x) // 10
		x := 5         // x being shadowed
		fmt.Println(x) // 5
	}

	// A shadowing variable is a variable that has the same name as a variable in a containing block. For as long as the shadowing variable exists, you cannot access a shadowed variable.

	fmt.Println(x) // 10

	fmt.Println("")

}
