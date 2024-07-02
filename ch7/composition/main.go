package main

import "fmt"

func main() {
	callManager()
	fieldWithSameName()
	notInheritance()
	noDynamicDIspatch()
}

// Furthermore, Go has no dynamic dispatch for concrete types. The methods on the embedded field have no idea they are embedded. If you have a method on an embedded field that calls another method on the embedded field, and the containing struct has a method of the same name, the method on the embedded field is invoked, not the method on the containing struct.

type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}
func (i Inner) Double() string {
	return i.IntPrinter(i.A * 2)
}

type Outer struct {
	Inner
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}
func noDynamicDIspatch() {
	o := Outer{
		Inner: Inner{
			A: 10},
		S: "Hello"}
	fmt.Println(o.Double()) // Inner: 20
}

// =============================
// Embedding Is Not Inheritance
// You cannot assign a variable of type Manager to a variable of type Employee. If you want to access the Employee field in Manager, you must do so explicitly

func notInheritance() {
	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}

	// var eFail Employee = m        // compilation error!
	var eOK Employee = m.Employee // ok!

	fmt.Println(eOK.ID)          // print 12345
	fmt.Println(m.Description()) // prints Bob Bobson (12345)
}

// =============================
// Use Embedding for Composition
// The software engineering advice “Favor object composition over class inheritance” dates back to at least the 1994 book Design Patterns by Erich Gamma, Richard Helm, Ralph Johnson, and John Vlissides (Addison-Wesley), better known as the Gang of Four book. While Go doesn’t have inheritance, it encourages code reuse via built-in support for composition and promotion:

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

// func (m Manager) FindNewEmployees() []Employee {

// 	// do business logic
// }

// Note that Manager contains a field of type Employee, but no name is assigned to that field. This makes Employee an embedded field. Any fields or methods declared on an embedded field are promoted to the containing struct and can be invoked directly on it. That makes the following code valid:
func callManager() {
	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)            // print 12345
	fmt.Println(m.Description()) // prints Bob Bobson (12345)
}

// You can embed any type within a struct, not just another struct. This promotes the methods on the embedded type to the contain‐ ing struct.

// If the containing struct has fields or methods with the same name as an embedded field, you need to use the embedded field’s type to refer to the obscured fields or methods. If you have types defined like this:

func fieldWithSameName() {
	type Inner struct {
		X int
	}
	type Outer struct {
		Inner
		X int
	}
	// you can access the X on Inner only by specifying Inner explicitly:
	o := Outer{
		Inner: Inner{
			X: 10},
		X: 20}
	fmt.Println(o.X)       // print 20
	fmt.Println(o.Inner.X) // prints 10
}

// =============================
