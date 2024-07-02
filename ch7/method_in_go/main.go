package main

import (
	"fmt"
	"time"
)

func main() {
	methodInvocation()
	usingReceiver()
	passingValueToReceiver()
	usingTree()
	callAdder()
}

// =============================

// Type Declarations Aren’t Inheritance
// In addition to declaring types based on built-in Go types and struct literals, you can also declare a user-defined type based on another user-defined type:

type Score int

type HighScore Score
type Employee Person

// That’s not the case in Go. You can’t assign an instance of type HighScore to a variable of type Score, or vice versa, without a type conversion, nor can you assign either of them to a variable of type int without a type conversion. Furthermore, any methods defined on Score aren’t defined on HighScore:

// assigning untyped constants is valid
func assigningUntyped() {
	var i int = 300
	var s Score = 100
	var hs HighScore = 200

	// hs = s            // compilation error!
	// s = i             // compilation error!
	s = Score(i)      // ok
	hs = HighScore(s) // ok

	fmt.Println(hs) // prints 15

}

// =============================

// Functions Versus Methods
// Since you can use a method as a function, you might wonder when you should declare a function and when you should use a method.
// The differentiator is whether your function depends on other data. As I’ve covered several times, package-level state should be effectively immutable. Anytime your logic depends on values that are configured at startup or changed while your program is running, those values should be stored in a struct, and that logic should be imple‐ mented as a method. If your logic depends only on the input parameters, it should be a function.

// ==========================

// Methods Are Functions Too

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}
func callAdder() {
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5)) // prints 15

	// You can also assign the method to a variable or pass it to a parameter of type func(int)int. This is called a method value:
	f1 := myAdder.AddTo
	fmt.Println(f1(10)) // prints 20

	// A method value is a bit like a closure, since it can access the values in the fields of the instance from which it was created.
	// You can also create a function from the type itself. This is called a method expression:

	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 15)) // prints 25

	// In a method expression, the first parameter is the receiver for the method; the
	// function signature is func(Adder, int) int.
	// Method values and method expressions aren’t clever corner cases. You’ll see one way to use them when you look at dependency injection in “Implicit Interfaces Make Dependency Injection Easier” on page 174.
}

//=============================

// Code Your Methods for nil Instances

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}
func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

func usingTree() {
	fmt.Println("Using Tree")
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))  // true
	fmt.Println(it.Contains(12)) // false
	fmt.Println(it.Contains(3))  // true
	fmt.Println(it)
	fmt.Println(it.left)
	fmt.Println(it.left.left)
	fmt.Println("")

}

//==============================

// Passing Value to receiver

// Be aware that the rules for passing values to functions still apply. If you pass a value type to a function and call a pointer receiver method on the passed value, you are invoking the method on a copy.

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}
func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}

func passingValueToReceiver() {
	fmt.Println("Passing Value to receiver")
	var c Counter
	doUpdateWrong(c) // passing a copy
	fmt.Println("in main:", c.String())
	doUpdateRight(&c)
	fmt.Println("in main:", c.String())
	fmt.Println("")

}

//==================================

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
	fmt.Println("Using Value Receiver")
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
	fmt.Println("")

	// If you call a value receiver on a pointer variable, Go automatically dereferences the pointer when calling the method. In the code:

	fmt.Println("Using Pointer Receiver")
	d := &Counter{}
	fmt.Println(d.String())
	d.Increment()
	fmt.Println(d.String())
	fmt.Println("")

	// the call c.String() is silently converted to (*c).String().

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
