package main

import "fmt"

func main() {
	pointerTo()
	dereference()
	panicPointer()
	wontCompilePointer()

	tryAndFail()
	tryAndSucceed()
}

// ref deref

func failedUpdate(g *int) {
	x := 10
	g = &x
}

func tryAndFail() {
	fmt.Println("try and failed")

	var f *int // f is nil
	failedUpdate(f)
	fmt.Println(f) // prints nil
	fmt.Println("")

}

func sndFailedUpdate(px *int) {
	x2 := 20
	px = &x2
}
func update(px *int) {
	*px = 20
}
func tryAndSucceed() {
	fmt.Println("Try And Succeed")

	x := 10
	sndFailedUpdate(&x)
	fmt.Println(x) // prints 10
	update(&x)
	fmt.Println(x) // prints 20
	fmt.Println("")

}

// You can’t use an & before a primitive literal (numbers, booleans, and strings) or a constant because they don’t have memory addresses; they exist only at compile time.

func wontCompilePointer() {
	fmt.Println("won't compile pointer")

	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}
	// decomment this to see the compiler complaining
	p := person{
		FirstName: "Pat",
		// MiddleName: "Perry", // This line won't compile
		LastName: "Peterson",
	}
	fmt.Println(p)

	// There are two ways around this problem. The first is to do what was shown previ‐ ously, which is to introduce a variable to hold the constant value. The second way is to write a generic helper function that takes in a parameter of any type and returns a pointer to that type:
	// see makePointer func

	// With that function, you can now write:
	p = person{
		FirstName:  "Pat",
		MiddleName: makePointer("Perry"), // This works
		LastName:   "Peterson",
	}

	fmt.Println(p)
	fmt.Println(*p.MiddleName)

	fmt.Println("")

}

func makePointer[T any](t T) *T {
	return &t
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

func panicPointer() {
	fmt.Println("Panic pointer")

	// Before dereferencing a pointer, you must make sure that the pointer is non-nil. Your program will panic if you attempt to dereference a nil pointer:
	var x *int
	fmt.Println(x == nil) // prints true
	// fmt.Println(*x)       // panic

	// The built-in function new creates a pointer variable. It returns a pointer to a zero- value instance of the provided type:
	x = new(int)
	fmt.Println(x == nil) // prints false
	fmt.Println(*x)       // prints 0

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
