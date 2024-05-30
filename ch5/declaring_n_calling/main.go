package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	callDivN()
	callingMYFunc()
	callingAddTo()
	callingDivAnd()
	callingNamedDivAnd()
	callingBlankDivAnd()
	callingMyFuncVariable()
	calculator()
	calculatorWithOpMapType()
	anonymFn()
	anonymFnDirectly()
	callAdd1()
	callingChangeAdd1()
}

// Since you can declare variables at the package scope, you can also declare package scope variables that are assigned anonymous functions:

var (
	add1 = func(i, j int) int { return i + j }
	sub1 = func(i, j int) int { return i - j }
	mul1 = func(i, j int) int { return i * j }
	div1 = func(i, j int) int { return i / j }
)

func callAdd1() {
	x := add1(2, 3)
	fmt.Println(x)
}

func callingChangeAdd1() {
	x := add1(2, 3)
	fmt.Println(x)
	changeAdd1()
	y := add1(2, 3)
	fmt.Println(y)
}
func changeAdd1() {
	add1 = func(i, j int) int { return i + j + j }
}

// You don’t have to assign an anonymous function to a variable. You can write them inline and call them immediately. The previous program can be rewritten into this:

func anonymFnDirectly() {
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside of an anonymous function")
		}(i)
	}
}

// Anonymous Functions
// You can not only assign functions to variables, but also define new functions within a function and assign them to variables. Here’s a program to demonstrate this, which you can run on The Go Playground. The code is also available in the sample_code/ anon_func directory in the Chapter 5 repository:

func anonymFn() {
	f := func(j int) { // error if giving name in this anonymous fn.
		fmt.Println("printing", j, "from inside of an anonymous function")
	}
	for i := 0; i < 5; i++ {
		f(i)
	}
}

// Function Type Declarations
// Just as you can use the type keyword to define a struct, you can use it to define a function type too (I’ll go into more details on type declarations in Chapter 7):

type opFuncType func(int, int) int

// You can then rewrite the opMap declaration to look like this:
var opMapWithType = map[string]opFuncType{ // same as before
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func calculatorWithOpMapType() {

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMapWithType[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}

// Calculator

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func calculator() {

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}

// Functions Are Values
// Just as in many other languages, functions in Go are values. The type of a function is built out of the keyword func and the types of the parameters and return values. This combination is called the signature of the function. Any function that has the exact same number and types of parameters and return values meets the type signature.

var myFuncVariable func(string) int

// myFuncVariable can be assigned any function that has a single parameter of type string and returns a single value of type int. Here’s a longer example:

func f1(a string) int {
	return len(a)
}
func f2(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}
	return total
}
func callingMyFuncVariable() {

	myFuncVariable = f1
	result := myFuncVariable("Hello")
	fmt.Println(result)
	myFuncVariable = f2
	result = myFuncVariable("Hello")
	fmt.Println(result)
}

// Blank Returns—Never Use These!
// If you use named return values, you need to be aware of one severe misfeature in Go: blank (sometimes called naked) returns. If you have named return values, you can just write return without specifying the values that are returned. This returns the last values assigned to the named return values. Rewrite the divAndRemainder function one last time, this time using blank returns:

func callingBlankDivAnd() {
	result, remainder, err := blankDivAndRemainder(5, 3)
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
	fmt.Println(result, remainder)
}

func blankDivAndRemainder(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	result, remainder = num/denom, num%denom
	return
}

// Named Return Values
// In addition to letting you return more than one value from a function, Go allows you to specify names for your return values. Rewrite the divAndRemainder function one more time, this time using named return values:

func callingNamedDivAnd() {
	result, remainder, err := namedDivAndRemainder(5, 3)
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
	fmt.Println(result, remainder)
}

// One important thing to note: the name that’s used for a named returned value is local to the function; it doesn’t enforce any name outside of the function. It is perfectly legal to assign the return values to variables of different names:

func namedDivAndRemainder(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

// Multiple Return Values
// The first difference that you’ll see between Go and other languages is that Go allows for multiple return values. Let’s add a small feature to the previous division program. You’re going to return both the dividend and the remainder from the function. Here’s the updated function:

func callingDivAnd() {
	result, remainder, err := divAndRemainder(5, 3)
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
	fmt.Println(result, remainder)
}

func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

// Variadic Parameter
// Go supports variadic parameters. The variadic parameter must be the last (or only) parameter in the input parameter list. You indicate it with three dots (...) before the type.

func callingAddTo() {
	fmt.Println("calling variadic Parameter")
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3, 5}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))
	fmt.Println("")

}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// Named Arguments
// If you want to emulate named and optional parameters, define a struct that has fields that match the desired parameters, and pass the struct to your function. Example 5-1 shows a snippet of code that demonstrates this pattern.

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) error {
	// do something here
	return fmt.Errorf(fmt.Sprintf("<%s>", string(opts.FirstName)))
}

func callingMYFunc() {
	fmt.Println("Calling Named Prameter")

	error1 := MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      50,
	})
	fmt.Println("that is error 1:", error1)
	error2 := MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})
	fmt.Println("that is error 2:", error2)
	fmt.Println("")

}

func callDivN() {
	fmt.Println("divN")
	var n = divN(2, 1)
	fmt.Println("that is n:", n)
	fmt.Println("")
}
func divN(num int, denom int) int {
	if denom == 0 {
		return 0
	}
	return num / denom
}
