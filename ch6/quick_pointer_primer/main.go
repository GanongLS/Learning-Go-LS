package main

import "fmt"

func main() {
	pointerTo()
	dereference()
}

func dereference() {
	fmt.Println("dereference")

	// The * is the indirection operator. It precedes a variable of pointer type and returns the pointed-to value. This is called dereferencing:
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX) // prints a memory address fmt.Println(*pointerToX) // prints 10
	z := 5 + *pointerToX
	fmt.Println(z)

	fmt.Println("")

}

func pointerTo() {
	fmt.Println("Pointer to")

	// The & is the address operator. It precedes a value type and returns the address where the value is stored:
	x := "hello"
	pointerToX := &x
	fmt.Println(pointerToX)

	fmt.Println("")

}
