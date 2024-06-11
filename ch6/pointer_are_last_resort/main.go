package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	makeFoo()
	usingJSON()
}

// Using ref to marshal unmarshal json

// The only time you should use pointer parameters to modify a variable is when the function expects an interface. You see this pattern when working with JSON (I’ll talk more about the JSON support in Go’s standard library in “encoding/json” on page 327):

func usingJSON() {

	type Person struct { // Define a struct named Person
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var f Person
	err := json.Unmarshal([]byte(`{"name": "Bob", "age": 30}`), &f)
	fmt.Println(err)
}

func makeFoo() {
	fmt.Println("Make Foo with and without ref")
	var x Foo
	MakeFooDont(&x)
	var y, _ = MakeFooDo()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("")

}

// That said, you should be careful when using pointers in Go. As discussed earlier, they make it harder to understand data flow and can create extra work for the garbage collector. Rather than populating a struct by passing a pointer to it into a function, have the function instantiate and return the struct (see Examples 6-3 and 6-4).
// Example 6-3. Don’t do this

type Foo struct {
	Field1 string
	Field2 int
}

func MakeFooDont(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}

// Example 6-4. Do this
func MakeFooDo() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20}

	return f, nil
}
