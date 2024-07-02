package main

import "fmt"

func main() {
	printIota()
	intermittentlyIota()
}

// =========================

// The value of iota increments for each constant in the const block, whether or not iota is used to define the value of a constant. The following code demonstrates what happens when iota is used intermittently in a const block:
const (
	Field1 = 0        // iota = 0
	Field2 = 1 + iota // iota = 1
	Field3 = 20       // iota = 2
	Field4            // iota = 3
	Field5 = iota     // iota = 4
)

func intermittentlyIota() {
	fmt.Println(Field1, Field2, Field3, Field4, Field5)
}

// Field2 is assigned 2 because iota has a value of 1 on the second line in the const block. Field4 is assigned 20 because it has no type or value explicitly assigned, so it gets the value of the previous line with a type and assignment. Finally, Field5 gets the value 4 because it is the fifth line and iota starts counting from 0.

// This is the best advice I’ve seen on iota:
// Don’t use iota for defining constants where its values are explicitly defined (elsewhere). For example, when implementing parts of a specification and the specification says which values are assigned to which constants, you should explicitly write the constant values. Use iota for “internal” purposes only. That is, where the constants are referred to by name rather than by value. That way, you can optimally enjoy iota by inserting new constants at any moment in time / location in the list without the risk of breaking everything.
// —Danny van Heumen

// =========================

// When using iota, the best practice is to first define a type based on int that will represent all the valid values:
type MailCategory int

// Next, use a const block to define a set of values for your type:
const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)

func printIota() {
	fmt.Println(Uncategorized, Personal, Spam, Social, Advertisements)
}

// The first constant in the const block has the type specified, and its value is set to iota. Every subsequent line has neither the type nor a value assigned to it. When the Go compiler sees this, it repeats the type and the assignment to all the subsequent constants in the block, which is iota. The value of iota increments for each constant defined in the const block, starting with 0. This means that 0 is assigned to the first constant (Uncategorized), 1 to the second constant (Personal), and so on. When a new const block is created, iota is set back to 0.

//==============================
