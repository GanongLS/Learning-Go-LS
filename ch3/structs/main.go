package main

import (
	"fmt"
)

func main() {
	structsDeclaration()
	anonymousStruct()
	comparingStruct()
}

func comparingStruct() {
	fmt.Println("Comparing struct")

	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g = struct {
		name string
		age  int
	}{}
	// compiles -- can use = and == between identical named and anonymous structs
	g = f
	fmt.Println(f == g) // true

	type fourthPerson struct {
		firstName string
		age       int
	}

	type fifthPerson struct {
		name          string
		age           int
		favoriteColor string
	}

	var fourth = fourthPerson{firstName: "Ali", age: 30}
	var fifth = fifthPerson{name: "Ali", age: 30, favoriteColor: "blue"}
	fmt.Println(fourth)
	fmt.Println(fifth)

	// fmt.Println(fourth == fifth) // missmatch

	fmt.Println("")

}

func anonymousStruct() {

	fmt.Println("Anonymous Struct")

	var person struct {
		name string
		age  int
		pet  string
	}
	person.name = "bob"
	person.age = 50
	person.pet = "dog"
	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}

	fmt.Println(person)
	fmt.Println(pet)

	fmt.Println("")

}

type person struct {
	name string
	age  int
	pet  string
}

func structsDeclaration() {
	fmt.Println("Structs Declaration")

	// A struct type that’s defined within a function can be used only within that function. (You’ll learn more about functions in Chapter 5.)

	// Once a struct type is declared, you can define variables of that type:
	var fred person
	fmt.Println("fred:", fred)

	bob := person{}
	fmt.Println("bob name:", bob.name)

	// a struct literal can be specified as a comma-separated list of values for the fields inside of braces:
	julia := person{
		"Julia",
		40,
		"cat",
	}

	fmt.Println("julia:", julia)
	fmt.Println("julia name:", julia.name)

	// When using this struct literal format, a value for every field in the struct must be specified, and the values are assigned to the fields in the order they were declared in the struct definition.

	// The second struct literal style looks like the map literal style:

	beth := person{
		age:  30,
		name: "Beth",
	}
	fmt.Println("beth:", beth)
	fmt.Println("beth name:", beth.name)

	// You use the names of the fields in the struct to specify the values. This style has some advantages. It allows you to specify the fields in any order, and you don’t need to provide a value for all fields. Any field not specified is set to its zero value.

	// A field in a struct is accessed with dot notation:
	bob.name = "Bob"
	fmt.Println("bob name:", bob.name)

	fmt.Println("")

}
