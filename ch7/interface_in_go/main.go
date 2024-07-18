package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	usingInterface()
	callLogic()
	nilInterface()
	callInterfaceComparer()
	emptyInterface()
	emptyJsonInterface()
	assertionInterface()
	doThings(nil)
}

//========================

// Use Type Assertions and Type Switches Sparingly

// Type switch statements provide the ability to differentiate between multiple imple‐ mentations of an interface that require different processing. They are most useful when only certain possible valid types can be supplied for an interface. Be sure to include a default case in the type switch to handle implementations that aren’t known at development time. This protects you if you forget to update your type switch statements when adding new interface implementations:

// type Node interface {
// }

// type number int

// type operator string

// func (op *operator) process(a, b operator) {

// }

// type treeNode struct {
// 	val            Node
// 	lchild, rchild *treeNode
// }

// func walkTree(t *treeNode) (Node, error) {
// 	switch val := t.val.(type) {
// 	case nil:
// 		return 0, errors.New("invalid expression")
// 	case number:
// 		// we know that t.val is of type number, so return the // int value
// 		return int(val), nil
// 	case operator:
// 		// we know that t.val is of type operator, so
// 		// find the values of the left and right children, then // call the process() method on operator to return the // result of processing their values.
// 		left, err := walkTree(t.lchild)
// 		if err != nil {
// 			return 0, err
// 		}
// 		right, err := walkTree(t.rchild)
// 		if err != nil {
// 			return 0, err
// 		}
// 		return val.process(left, right), nil
// 	default:
// 		// if a new treeVal type is defined, but walkTree wasn't updated // to process it, this detects it
// 		return 0, errors.New("unknown node type")
// 	}
// }

//========================

// Type Assertions and Type Switches
// Go provides two ways to see if a variable of an interface type has a specific concrete type or if the concrete type implements another interface. Let’s start by looking at type assertions. A type assertion names the concrete type that implemented the interface, or names another interface that is also implemented by the concrete type whose value is stored in the interface.

type MyInt int

func assertionInterface() (err error) {
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2 + 1)

	// i2 = i.(string) // panic: interface conversion: interface {} is main.MyInt, not string
	// fmt.Println(i2)

	// i2 = i.(int)  // cannot use i.(int) (comma, ok expression of type int) as MyInt value in assignment
	// fmt.Println(i2 + 1)

	i3, ok := i.(int)
	if !ok {
		return fmt.Errorf("unexpected type for %v", i)
	}
	fmt.Println(i3 + 1)
	return nil
}

// Even if you are absolutely certain that your type assertion is valid, use the comma ok idiom version. You don’t know how other people (or you in six months) will reuse your code. Sooner or later, your unvalidated type assertions will fail at runtime.
// When an interface could be one of multiple possible types, use a type switch instead:
func doThings(i any) {
	switch j := i.(type) {
	case nil:
		// i is nil, type of j is any
		fmt.Println("i is nil, type of j is any", j)

	case int:
	// j is of type int
	case MyInt:
	// j is of type MyInt
	case io.Reader:
	// j is of type io.Reader
	case string:
		// j is a string
		fmt.Println("i is string, type of j is string", j)

	case bool, rune:
	// i is either a bool or rune, so j is of type any
	default:
		// no idea what i is, so j is of type any
	}
}

//========================

// The Empty Interface Says Nothing

// Sometimes in a statically typed language, you need a way to say that a variable could store a value of any type. Go uses an empty interface, interface{}, to represent this:

func emptyInterface() {

	fmt.Println("Empty Interface")

	var i interface{}
	i = 20
	i = "hello"
	i = struct {
		FirstName string
		LastName  string
	}{"Fred", "Fredson"}

	fmt.Println(i)

	fmt.Println("")

}

// Because an empty interface doesn’t tell you anything about the value it represents, you can’t do a lot with it. One common use of any is as a placeholder for data of uncertain schema that’s read from an external source, like a JSON file:

type Card interface{}

func emptyJsonInterface() (err error) {
	data := map[string]any{}
	contents, err := os.ReadFile("testdata/sample.json")
	if err != nil {
		return err
	}
	json.Unmarshal(contents, &data)
	// the contents are now in the data map
	return nil
}

//========================
// Interfaces Are Comparable

// In Chapter 3, you learned about comparable types, the ones that can be checked for equality with ==. You might be surprised to learn that interfaces are comparable. Just as an interface is equal to nil only if its type and value fields are both nil, two instances of an interface type are equal only if their types are equal and their values are equal. This suggests a question: what happens if the type isn’t comparable? Let’s use a simple example to explore this concept. Start with an interface definition and a couple of implementations of that interface:

type Doubler interface {
	Double()
}
type DoubleInt int

func (d *DoubleInt) Double() {
	*d = *d * 2
}

type DoubleIntSlice []int

func (d DoubleIntSlice) Double() {
	for i := range d {
		d[i] = d[i] * 2
	}
}

// The Double method on DoubleInt is declared with a pointer receiver because you are modifying the value of the int. You can use a value receiver for the Double method on DoubleIntSlice because, as covered in “The Difference Between Maps and Slices” on page 131, you can change the value of an item in a parameter that is a slice type. The *DoubleInt type is comparable (all pointer types are), and the DoubleIntSlice type is not comparable (slices aren’t comparable).
// You also have a function that takes in two parameters of type Doubler and prints if they are equal:

func DoublerCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

// You now define four variables:
var di DoubleInt = 10
var di2 DoubleInt = 10
var dis = DoubleIntSlice{1, 2, 3}
var dis2 = DoubleIntSlice{1, 2, 3}

// Now, you’re going to call this function three times. The first call is as follows:

func callInterfaceComparer() {
	fmt.Println("Interface Comparer")

	DoublerCompare(&di, &di2) // false

	// This prints out false. The types match (both are *DoubleInt), but you are compar‐
	// ing pointers, not values, and the pointers point to different instances. Next, you compare a *DoubleInt with a DoubleIntSlice:

	DoublerCompare(&di, dis)

	// This prints out false, because the types do not match.

	// Finally, you get to the problematic case:

	// DoublerCompare(dis, dis2) // run time error.

	// This code compiles without issue, but triggers a panic at runtime:

	fmt.Println(dis2)

	fmt.Println("")

}

//========================

// Accept Interfaces, Return Structs

// You’ll often hear experienced Go developers say that your code should “Accept inter‐ faces, return structs.” This phrase was most likely coined by Jack Lindamood in his 2016 blog post “Preemptive Interface Anti-Pattern in Go”. It means that the business logic invoked by your functions should be invoked via interfaces, but the output of your functions should be a concrete type. I’ve already covered why functions should accept interfaces: they make your code more flexible and explicitly declare the exact functionality being used.

// In order for an interface to be considered nil, both the type and the value must be nil. The following code prints out true on the first two lines and false on the last:
func nilInterface() {
	var pointerCounter *Counter
	fmt.Println(pointerCounter == nil) // prints true
	var incrementer Incrementer
	fmt.Println(incrementer == nil) // prints true
	incrementer = pointerCounter
	// fmt.Println(incrementer == nil) // prints false // always false
}

//========================

// Embedding and Interfaces

// Embedding is not only for structs. You can also embed an interface in an interface. For example, the io.ReadCloser interface is built out of an io.Reader and an io.Closer:
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}
type ReadCloser interface {
	Reader
	Closer
}

//========================

// Interfaces Are Type-Safe Duck Typing

// So far, nothing that’s been said about the Go interface is much different from inter‐ faces in other languages. What makes Go’s interfaces special is that they are imple‐ mented implicitly. As you’ve seen with the Counter struct type and the Incrementer interface type that you’ve used in previous examples, a concrete type does not declare that it implements an interface. If the method set for a concrete type contains all the methods in the method set for an interface, the concrete type implements the interface. Therefore, that the concrete type can be assigned to a variable or field declared to be of the type of the interface.

// This implicit behavior makes interfaces the most interesting thing about types in Go, because they enable both type safety and decoupling, bridging the functionality in both static and dynamic languages.

// Dynamically typed languages like Python, Ruby, and JavaScript don’t have interfaces. Instead, those developers use duck typing, which is based on the expression “If it walks like a duck and quacks like a duck, it’s a duck.” The concept is that you can pass an instance of a type as a parameter to a function as long as the function can find a method to invoke that it expects:

// Java developers use a different pattern. They define an interface, create an implemen‐ tation of the interface, but refer to the interface only in the client code:

// Dynamic language developers look at the explicit interfaces in Java and don’t see how you can possibly refactor your code over time when you have explicit dependencies. Switching to a new implementation from a different provider means rewriting your code to depend on a new interface.

// Go’s developers decided that both groups are right. If your application is going to grow and change over time, you need flexibility to change implementation. However, in order for people to understand what your code is doing (as new people work on the same code over time), you also need to specify what the code depends on. That’s where implicit interfaces come in. Go code is a blend of the previous two styles:

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string { // business logic
	return data
}

type Logic interface {
	Process(data string) string
}
type Client struct {
	L Logic
}

func (c Client) Program() {
	// get data from somewhere c.L.Process(data)
}
func callLogic() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}

// Interfaces specify what callers need. The client code defines the interface to specify what functionality it requires.

//========================

type Stringer interface {
	String() string
}

// In an interface declaration, an interface literal appears after the name of the interface type. It lists the methods that must be implemented by a concrete type to meet the interface. The methods defined by an interface are the method set of the interface. As I covered in “Pointer Receivers and Value Receivers” on page 145, the method set of a pointer instance contains the methods defined with both pointer and value receivers, while the method set of a value instance contains only the methods with value receivers. Here’s a quick example using the Counter struct defined previously:
type Incrementer interface {
	Increment()
}

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

func usingInterface() {
	var myStringer fmt.Stringer
	var myIncrementer Incrementer
	pointerCounter := &Counter{}
	valueCounter := Counter{}
	myStringer = pointerCounter // ok
	myStringer = valueCounter   // ok

	myIncrementer = pointerCounter // ok
	// myIncrementer = valueCounter   // compile-time error!

	fmt.Println(myStringer)
	fmt.Println(myIncrementer)
	fmt.Println(pointerCounter)
	fmt.Println(pointerCounter)

}

// Interfaces are usually named with “er” endings. You’ve already seen fmt.Stringer, but there are many more, including io.Reader, io.Closer, io.ReadCloser, json.Marshaler, and http.Handler.
