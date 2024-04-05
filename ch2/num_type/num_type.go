package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	complexExample()
	numTypeConversion()
	intTypeConv()
	unusedVar()
}

func unusedVar() {
	x := 10 // this assignment isn't read! 
	x = 20
	fmt.Println(x)
	x = 30 // this assignment isn't read!
}

func intTypeConv() {
	// Example 2-3. Integer type conversions
	var x int = 10
	var b byte = 100
	var sum3 int = x + int(b)
	var sum4 byte = byte(x) + b
	fmt.Println(sum3, sum4)

}

func numTypeConversion() {
	// Example 2-2. Type conversions
	var x int = 10
	var y float64 = 30.2
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)
	fmt.Println(sum1, sum2)

}

func complexExample() {
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))
}
