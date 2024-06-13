package main

import (
	"fmt"
	"time"
)

func main() {

	methodInvocation()
	usingReceiver()
}

// Pointer Receivers and Value Receivers

// • If your method modifies the receiver, you must use a pointer receiver.
// • If your method needs to handle nil instances (see “Code Your Methods for nil
// Instances” on page 148), then it must use a pointer receiver.
// • If your method doesn’t modify the receiver, you can use a value receiver.

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func usingReceiver() {
	fmt.Println("Using Receiver")
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
	fmt.Println("")

}

// ========================

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

// Method declarations look like function declarations, with one addition: the receiver specification. The receiver appears between the keyword func and the name of the method. Like all other variable declarations, the receiver name appears before the type. By convention, the receiver name is a short abbreviation of the type’s name, usually its first letter. It is nonidiomatic to use this or self.

// I’ll talk more about packages in Chapter 10, but be aware that methods must be declared in the same package as their associated type; Go doesn’t allow you to add methods to types you don’t control. While you can define a method in a different file within the same package as the type declaration, it is best to keep your type definition and its associated methods together so that it’s easy to follow the implementation.

// method invocation

func (p Person) getLastName() string {
	return p.LastName
}
func methodInvocation() {
	// Method invocations should look familiar to those who have used methods in other languages:
	p := Person{
		FirstName: "Fred",
		LastName:  "Fredson",
		Age:       52,
	}

	output := p.getLastName()
	sOutput := p.String()
	fmt.Println(output)
	fmt.Println(sOutput)
}
